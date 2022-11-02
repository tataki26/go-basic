package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h := hotels{
		{
			Name:    "hotel A",
			Address: "qwer",
			City:    "city A",
			Zip:     "1234",
			Region:  "Northern",
		},
		{
			Name:    "hotel B",
			Address: "tyui",
			City:    "city B",
			Zip:     "4567",
			Region:  "Central",
		},
		{
			Name:    "hotel C",
			Address: "asdf",
			City:    "city C",
			Zip:     "4275",
			Region:  "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}
