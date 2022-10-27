package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	fruits := []string{"Apple", "Banana", "Orange", "Mango", "Watermelon"}

	err := tpl.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatalln(err)
	}
}
