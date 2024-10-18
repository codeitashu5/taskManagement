package task

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RegisterBody struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

type Task struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Task       string             `json:"task" bson:"task"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	ArchivedAt *time.Time         `json:"archived_at,omitempty" bson:"archived_at"`
}

type Request struct {
	Task string `json:"task" bson:"task"`
}
