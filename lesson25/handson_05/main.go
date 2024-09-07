package main

import (
	"context"
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 007)
	ctx = context.WithValue(ctx, "fname", "James")

	result := dbAcces(ctx)
	fmt.Fprintln(w, result)
}

func dbAcces(ctx context.Context) int {
	uid := ctx.Value("userID").(int)
	return uid
}
