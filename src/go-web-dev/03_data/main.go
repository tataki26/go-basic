package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// initalize package
// when package is loaded, init function is called first
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// use data argument
	// send it to gohtml file
	// only one data(or data structure)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
