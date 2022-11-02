package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := Regions{
		{
			Region: "Southern",
			Hotels: []hotel{
				{
					Name:    "Hotel A",
					Address: "qwer",
					City:    "City A",
					Zip:     "1234",
					Region:  "southern",
				},
			},
		},
		{
			Region: "Central",
			Hotels: []hotel{
				{
					Name:    "Hotel B",
					Address: "oipe",
					City:    "City B",
					Zip:     "7894",
					Region:  "central",
				},
			},
		},
		{
			Region: "Northern",
			Hotels: []hotel{
				{
					Name:    "Hotel C",
					Address: "zcdf",
					City:    "City C",
					Zip:     "8546",
					Region:  "northern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
