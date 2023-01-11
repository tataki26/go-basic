package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// BSON: binary JSON
// ObjectID: Primary Key(unique - RDBMS)
type User struct {
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
	Id     primitive.ObjectID `json:"id" bson:"_id"` // string before
}
