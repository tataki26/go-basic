package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {
	var id int
	var err error

	if req.Method == http.MethodPost {
		id, err = strconv.Atoi(req.FormValue("ethernetid"))
		if err != nil {
			log.Fatal(err)
		}
	}

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Do("SET", "id", id)
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", id)
}
