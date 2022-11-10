package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	// need only ResponseWriter, *Request, and file name
	http.ServeFile(w, req, "dog.jpg")
}
