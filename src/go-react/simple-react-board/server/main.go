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

	mux.HandleFunc("/api/get", get)
	mux.HandleFunc("/api/insert", insert)
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
}

var db *sql.DB
var err error

type User struct {
	Title   string
	Content string
}

type Data struct {
	Index   string
	Title   string
	Content string
}

// GET: 페이지가 랜더링될 때마다 DB에 저장된 data 가져와서 Form에 전달
func get(w http.ResponseWriter, req *http.Request) {

	datas := []*Data{}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/simpleboard")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := db.Query("SELECT * FROM simpleboard;")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	for conn.Next() {
		var data Data
		conn.Scan(&data.Index, &data.Title, &data.Content)

		datas = append(datas, &data)
	}

	json.NewEncoder(w).Encode(datas)

	defer db.Close()

}

// POST: 버튼 클릭 시, input data json으로 받아와 DB에 저장
func insert(w http.ResponseWriter, req *http.Request) {
	var u User

	err = json.NewDecoder(req.Body).Decode(&u)
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
