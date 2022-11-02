package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := region{
		Region: "Southern",
		Hotels: []hotel{
			{
				Name:    "Hotel A",
				Address: "qwer",
				City:    "City A",
				Zip:     "1234",
			},
			{
				Name:    "Hotel B",
				Address: "uiop",
				City:    "City B",
				Zip:     "7896",
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
