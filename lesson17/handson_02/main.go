package main

import "net/http"

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
