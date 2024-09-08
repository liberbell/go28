package main

import (
	"context"
)

func main() {
	ctx, channel := context.WithCancel(context.Background())
	defer channel()
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
