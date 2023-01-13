package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectID: Primary Key(unique - RDBMS)
type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}
