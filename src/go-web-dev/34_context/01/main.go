package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// return request context
	// to control work flow(set work time, work cancellation, ...)
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
