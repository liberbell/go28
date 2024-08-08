package main

import "net/http"

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Original-key", "This is an original key")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
