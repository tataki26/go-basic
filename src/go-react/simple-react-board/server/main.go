package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("server is running...")

	// open mysql server
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/simpleboard")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("DB 연동: %+v\n", db.Stats())

	// execute query
	conn, err := db.Query("INSERT INTO simpleboard (title, content) VALUES ('test', 'value')")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Connection 생성: %+v\n", db.Stats())
	io.WriteString(w, "success!!")

	/*
		// bring data from database
		for conn.Next() {
			idx := ""
			name := ""
			value := ""
			if err := conn.Scan(&idx, &name, &value); err != nil {
				fmt.Println(err)
			}
			fmt.Println(idx, name, value)
		}
	*/

	conn.Close()
	fmt.Printf("Connection 연결 종료: %+v\n", db.Stats())

	db.Close()
	fmt.Printf("DB 연동 종료: %+v\n", db.Stats())
}
