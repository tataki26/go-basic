package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// package "template" has "ParseFiles" method
	// ParseFiles has multiple string arguments(file name) and returns tpl(pointer)
	// think pointer as container which has all parsed templates
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute creates writer interface(from os.Stdout) and data - arguments
	// os.Stdout for print
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
