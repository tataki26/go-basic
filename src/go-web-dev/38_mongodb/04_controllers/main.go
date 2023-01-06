package main

import (
	"mongodb/04_controllers/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// return UserController pointer
	// to use its methods
	uc := controllers.NewUserController()

	// no index - use curl
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
