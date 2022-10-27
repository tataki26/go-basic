package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type fruit struct {
	Name  string
	Color string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// composite data structures
	// struct
	watermelon := fruit{
		Name:  "Watermelon",
		Color: "Green",
	}

	err := tpl.Execute(os.Stdout, watermelon)
	if err != nil {
		log.Fatalln(err)
	}
}
