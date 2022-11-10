package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	// set response header(in develop tool under the network tab)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// image can be posted by using io.WriteString(HTML tag)
	io.WriteString(w, `
	<!-- not serving from our server -->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
