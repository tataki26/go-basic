package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// two different handler
	var d hotdog
	var c hotcat

	// routing
	mux := http.NewServeMux()
	// "/dog/": when you put another words behind it, it also calls dog code
	// ex:) "/dog/something/else"
	mux.Handle("/dog/", d)
	// exact match
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)

}
