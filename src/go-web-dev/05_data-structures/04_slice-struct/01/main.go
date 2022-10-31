package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// struct
type fruit struct {
	Name  string
	Color string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Food      []fruit
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	a := fruit{
		Name:  "Apple",
		Color: "Red",
	}

	b := fruit{
		Name:  "Banana",
		Color: "Yellow",
	}

	w := fruit{
		Name:  "Watermelon",
		Color: "Green",
	}

	f := car{
		Manufacturer: "Ford",
		Doors:        2,
	}

	v := car{
		Manufacturer: "Volkswagen",
		Doors:        4,
	}

	// slice
	fruits := []fruit{a, b, w}
	cars := []car{f, v}

	data := items{
		Food:      fruits,
		Transport: cars,
	}

	// type of data: struct
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
