package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodb/05_mongodb/05_update-user-controllers-delete/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	u := models.User{}

	cu := uc.client.Database("go-web-dev-db").Collection("users")
	err = cu.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID()

	_, err := uc.client.Database("go-web-dev-db").Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		log.Println(err)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// delete
	cu := uc.client.Database("go-web-dev-db").Collection("users")
	_, err = cu.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
