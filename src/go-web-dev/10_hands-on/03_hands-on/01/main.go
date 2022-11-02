package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip, Meal string
	Price               int64
}

type Items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	i := Items{
		{
			Name:    "Bread",
			Descrip: "Good",
			Meal:    "Breakfast",
			Price:   2500,
		},
		{
			Name:    "Pasta",
			Descrip: "Delicious",
			Meal:    "Lunch",
			Price:   8000,
		},
		{
			Name:    "Meat Soup",
			Descrip: "tasty",
			Meal:    "Dinner",
			Price:   12000,
		},
	}

	err := tpl.Execute(os.Stdout, i)
	if err != nil {
		log.Fatalln(err)
	}
}
