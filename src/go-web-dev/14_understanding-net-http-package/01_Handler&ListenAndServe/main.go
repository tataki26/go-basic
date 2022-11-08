package main

import (
	"fmt"
	"net/http"
)

// create my own type
// underlying type --> int
type hotdog int

// Handler interface
// the method of hotdog type
// ServeHTTP(ResponseWriter, *Request)
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// write to stream(w)
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var h hotdog
	// if something comes into ":8080", h(Handler) handles it
	// it has net.Listen(Listener) / Accept
	http.ListenAndServe(":8080", h)
}
