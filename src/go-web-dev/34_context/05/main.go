package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Background returns an empty Context which is never canceled, has no values, and has no deadline
	// WithCancel returns a copy of parent with a new Done channel
	// send context to new cancel and delay it
	ctx, cancel := context.WithCancel(context.Background())
	// make sure all paths cancel the context to avoid context leak(process keeps running <-- avoid)
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			// cancle to avoid process leak
			cancel()
			break
		}
	}

	time.Sleep(1 * time.Minute)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int

		for {
			select {
			case <-ctx.Done():
				return // avoid leaking of this goroutine when ctx is done
			case ch <- n: // keep increasing when it gets data
				n++
			}
		}
	}()

	return ch
}
