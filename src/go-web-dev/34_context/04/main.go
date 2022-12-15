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

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	// set timeout --> finish after 1 second
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// delay cancel --> WithTimeout returns CancelFunc instead of error
	defer cancel()

	// create int channel
	ch := make(chan int)

	go func() {
		// ridiculous long running task --> it occurs error
		uid := ctx.Value("userID").(int)
		// sleep 10 seconds
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// check error
		if ctx.Err() != nil {
			return
		}

		// send uid to channel
		ch <- uid
	}()

	// switch for channel
	// block until channel gets value
	select {
	// Done returns a channel that's closed when work done on behalf of this context should be canceled
	// if context canceled, make context done
	// wait until goroutine above finishes
	case <-ctx.Done(): // channel event
		return 0, ctx.Err() // timeout error
	// get value from channel
	case i := <-ch: // if the value is in the channel, it will be removed
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
