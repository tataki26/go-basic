package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// no request for favicon
	if req.URL.Path != "/" {
		// reply to the request with an HTTP 404 found error
		// it is the base of the NotFoundHandler method
		http.NotFound(w, req)
		return
	}

	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
