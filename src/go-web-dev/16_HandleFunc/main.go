package main

import (
	"io"
	"net/http"
)

// HandleFunc changes function into handler(func(ResponseWriter, *Request))
// not need ServeHTTP
// d is identifier
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// register the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}
