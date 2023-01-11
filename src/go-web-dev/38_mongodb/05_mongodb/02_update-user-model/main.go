package main

import (
	"context"
	"fmt"
	"mongodb/05_mongodb/02_update-user-model/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getClient())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getClient() *mongo.Client {
	// Connect to our local mongo
	// port: 27017
	// context: timeout, deadline, ...
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))
	// check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	fmt.Println("몽고 DB에 연결했습니다")

	return c
}
