package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// create new html file - nf
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("fail to create file", err)
	}
	defer nf.Close()

	// execute file without data
	// write result in nf
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
