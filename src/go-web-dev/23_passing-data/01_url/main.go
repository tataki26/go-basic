package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// request form value
	// q: identifier
	// if it exists, FormValue returns value(v)
	// return the first value of the query
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}
