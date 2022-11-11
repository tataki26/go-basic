package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// *** don't get confused ***
	// Handle(assets in url)
	// StripPrefix(assets in url)
	// Dir(assets in server directory)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// assets in url
	io.WriteString(w, `<img src="/assets/dog.jpg">`)
}
