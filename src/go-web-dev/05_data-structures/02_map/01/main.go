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
	// composite data structure
	// string map [key: value]
	fruits := map[string]string{
		"Apple":      "Red",
		"Banana":     "Yellow",
		"Orange":     "Orange",
		"Mango":      "Yellow",
		"Watermelon": "Green",
	}

	err := tpl.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatalln(err)
	}
}
