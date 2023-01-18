package main

import (
	"mongodb/08_hands-on/controllers"
	"mongodb/08_hands-on/models"
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
	// not create new map
	// load data from file - data permanence
	return models.LoadUsers()
}
