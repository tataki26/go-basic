package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"mongodb/05_mongodb/04_update-user-controllers-get/models"
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
	// Grab id(hex)
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// composite literal
	u := models.User{}

	// Fetch user
	// collection: group of Document (similar with table)
	cu := uc.client.Database("go-web-dev-db").Collection("user")

	// get a record and put into u
	if err := cu.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u); err != nil {
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

	// get user input
	// store it in User struct
	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID (update ID)
	u.Id = primitive.NewObjectID()

	// store the user in mongodb
	uc.client.Database("go-web-dev-db").Collection("users").InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
