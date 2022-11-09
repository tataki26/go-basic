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
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)

}
