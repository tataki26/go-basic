package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// server
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

// AJAX request
// - make browser **automatically** react to certain action
// - not to change whole page ---> change only part of page
// - use only required data
func foo(w http.ResponseWriter, r *http.Request) {
	s := `Here is some text from foo`

	// send string to browser
	fmt.Fprintln(w, s)
}
