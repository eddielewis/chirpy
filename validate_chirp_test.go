package main

import "testing"

// func TestCreateCache(t *testing.T) {
// 	interval := time.Second
// 	cache := NewCache(interval)
// 	if cache.cache == nil {
// 		t.Error("cache is nil")
// 	}
// }

// func TestAddGetCache(t *testing.T) {
// 	interval := time.Second
// 	cache := NewCache(interval)

// 	cache.Add("key1", []byte("val1"))
// 	actual, ok := cache.Get("key1")
// 	if !ok {
// 		t.Error("key1 not found")
// 	}
// 	if string(actual) != "val1" {
// 		t.Error("value doesn't match")
// 	}
// }

func TestReplaceProfanity(t *testing.T) {
	chirp := "I hear Mastodon is better than Chirpy. sharbert I need to migrate"
	out := removeProfanity(chirp)
	if out != "I hear Mastodon is better than Chirpy. **** I need to migrate" {
		t.Error("cleaned string does not match expected output")
	}
}
