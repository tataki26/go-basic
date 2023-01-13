package main

import (
	"mongodb/07_solution/controllers"
	"mongodb/07_solution/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getClient())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getClient() map[string]models.User {
	// make client
	return make(map[string]models.User)
}
