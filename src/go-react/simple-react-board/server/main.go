package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/insert", insert)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
}

var db *sql.DB

type User struct {
	Title   string
	Content string
}

// POST: 버튼 클릭 시, input data json으로 받아와 DB에 저장
func insert(w http.ResponseWriter, req *http.Request) {
	var u User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := u.Title
	content := u.Content

	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/simpleboard")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := db.Query("INSERT INTO simpleboard (title, content) VALUES (?,?)", title, content)
	if err != nil {
		log.Fatalln(err)
	}

	conn.Close()
	fmt.Printf("Connection 연결 종료: %+v\n", db.Stats())

	db.Close()
	fmt.Printf("DB 연동 종료: %+v\n", db.Stats())
}
