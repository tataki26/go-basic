package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is index page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "i love dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "takityaki")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
