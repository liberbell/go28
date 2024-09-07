package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "James")

	result := dbAcces(ctx)
	fmt.Fprintln(w, result)
}

func dbAcces(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 1)
	defer cancel()

	ch := make(chan, int)

	go func ()  {
		uid := ctx.Value("userID").(int)
		time.Sleep(time.Second * 2)

		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
