package main

import (
	"io"
	"net/http"
)

// ordinary functions --> not Handler type
// it is underlying type of HandlerFunc
// cause it has ResponseWriter and *Request
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
