package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/cors"
)

type Connection struct {
	Id string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/insert", insert)
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
}

func insert(w http.ResponseWriter, req *http.Request) {
	var c Connection

	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := c.Id

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Do("SET", "id", id)
	if err != nil {
		log.Fatal(err)
	}

}
