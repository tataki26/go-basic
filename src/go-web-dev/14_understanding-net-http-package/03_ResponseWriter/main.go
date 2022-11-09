package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// before we write header, we should set header configuration
	w.Header().Set("Mcleaod-Key", "this is from mcleaod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// write to ResponseWriter stream
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
