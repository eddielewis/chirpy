package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var profanities = [...]string{"kerfuffle", "sharbert", "fornax"}

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Error       string `json:"error"`
		CleanedBody string `json:"valid"`
	}

	params := parameters{}
	respBody := returnVals{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		errorMsg := fmt.Sprintf("Error decoding parameters: %s", err)
		log.Print(errorMsg)
		respondWithError(w, 500, errorMsg)
		return
	}

	switch {
	case len(params.Body) < 1:
		respondWithError(w, 400, "Chirp is empty!")
		return
	case len(params.Body) > 140:
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	cleanedBody := removeProfanity(params.Body)
	respBody.CleanedBody = cleanedBody

	respondWithJSON(w, 200, respBody)
}

func removeProfanity(chirp string) string {
	mask := "****"
	words := strings.Split(chirp, " ")

	for x := 1; x < len(words); x++ {
		w := strings.ToLower(words[x])
		for y := 0; y < len(profanities); y++ {
			if profanities[y] == w {
				words[x] = mask
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(msg))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)

	dat, err := json.Marshal(payload)
	if err != nil {
		errorMsg := fmt.Sprintf("Error marshalling response JSON: %s", err)
		log.Print(errorMsg)
		respondWithError(w, 500, errorMsg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dat)
}
