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

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	r := restaurants{
		{
			Name: "rest A",
			Menu: menu{
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
			},
		},
		{
			Name: "rest B",
			Menu: menu{
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
			},
		},
		{
			Name: "rest C",
			Menu: menu{
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
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}

}
