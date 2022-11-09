package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// third party package
	// return new router
	mux := httprouter.New()

	// intuitive
	// use router's method which has same name with http method
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	// able to put variable into the name of the path
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/user/:category/:article", blogWrite)

	http.ListenAndServe(":8080", mux)
}

// get parameter from path name
func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
