package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	// _ : blank alias
	// never use its method, use only for driver setup
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// driver & data source name --> string
	// return db pointer ---> able to use its method
	// sql --> from standard sql library
	db, err = sql.Open("mysql", "root:0000@tcp(localhost:3306)/test02?charset=utf8")
	check(err)
	// necessary
	defer db.Close()

	// check if db works or not
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
