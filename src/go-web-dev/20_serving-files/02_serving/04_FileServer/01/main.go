package main

import (
	"io"
	"net/http"
)

func main() {
	// FileServer serves file server
	// Dir(.): current directory(every file in 01 directory)
	// --> file system / open method
	// return Handler
	// give browser "dog.jpg"
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// browser requests "dog.jpg"
	io.WriteString(w, `<img src="/dog.jpg">`)
}
