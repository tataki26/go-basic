package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type fruit struct {
	Name  string
	Color string
}

type car struct {
	Manufacturer string
	Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func[value] from package strings[key]
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// chaining
	// template.New: create a new template with some name
	// Funcs: add the elements of the argument map to the template's function map
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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

	fruits := []fruit{a, b, w}
	cars := []car{f, v}

	data := struct {
		Food      []fruit
		Transport []car
	}{
		fruits,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
