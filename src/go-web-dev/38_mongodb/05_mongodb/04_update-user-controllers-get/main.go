package main

import (
	"context"
	"fmt"
	"mongodb/05_mongodb/04_update-user-controllers-get/controllers"
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
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		panic(err)
	}

	fmt.Println("몽고 DB에 연결했습니다")

	return c
}
