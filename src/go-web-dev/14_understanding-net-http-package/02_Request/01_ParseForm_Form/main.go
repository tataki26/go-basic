package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// parse the raw query from the URL and update Form(the field of Request)
	// this method attaches to Request pointer
	// when request comes, the form is parsed
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// Form contains the parsed form data(URL, query parameters)
	// this field is only available after ParseForm is called
	// able to access data(from url, body) which is sent with request
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
