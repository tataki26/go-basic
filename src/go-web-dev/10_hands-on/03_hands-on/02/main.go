package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         int64
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := menu{
		{
			Meal: "Breakfast",
			Items: []item{
				{
					Name:    "Bread",
					Descrip: "Good",
					Price:   2500,
				},
			},
		},
		{
			Meal: "Lunch",
			Items: []item{
				{
					Name:    "Pasta",
					Descrip: "Delicious",
					Price:   8000,
				},
			},
		},
		{
			Meal: "Dinner",
			Items: []item{
				{
					Name:    "Meat Soup",
					Descrip: "tasty",
					Price:   12000,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}

}
