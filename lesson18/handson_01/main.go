package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}
