package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "do my search: "+v)
}
