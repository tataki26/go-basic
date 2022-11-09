package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is index page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "i love dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("me.gohtml")
	if err != nil {
		log.Fatalln("parsing error", err)
	}

	err = tpl.ExecuteTemplate(res, "me.gohtml", "takityaki")
	if err != nil {
		log.Fatalln("executing error", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
