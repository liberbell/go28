package main

import (
	"context"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 007)
	ctx = context.WithValue(ctx, "fname", "James")
}
