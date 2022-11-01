package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Age   int
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	u1 := user{
		Name:  "Tom",
		Age:   42,
		Admin: true,
	}

	u2 := user{
		Name:  "Sarah",
		Age:   29,
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Age:   5,
		Admin: false,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
