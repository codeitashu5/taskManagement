package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// The User model
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
