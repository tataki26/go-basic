package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mongodb/03_post_delete/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// added route
	// post - get posted data and create user
	r.POST("/user", createUser)
	// added route plus parameter
	r.DELETE("/user/:id", deleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
	</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// composite literal - type and curly braces
	// empty (to store data from user input)
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatalln(err)
	}

	// Change Id
	u.Id = "007"

	// marshal/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// not actual delete process (no store, no delete) - for practice
func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
