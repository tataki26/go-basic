package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// StripPrefix returns handler
	// include all sub direcories of resources directory
	// anything that comes into resources(url), use that handler
	// remove given prefix(from url)
	// result ---> ./assets/dog.jpg
	// you can serve only jpg file ---> you don't neeed to serve the code
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// remove resources(prefix) --> "dog.jpg"
	io.WriteString(w, `<img src="/resources/dog.jpg>`)
}
