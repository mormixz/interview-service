package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interview struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Description   string             `json:"description" bson:"description"`
	Status        string             `json:"status" bson:"status"`
	CreatedBy     primitive.ObjectID `json:"-" bson:"created_by"`
	CreatedByUser Users              `json:"created_by" bson:"-"`
	CreatedAt     *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
	Comments      []Comment          `json:"comments,omitempty" bson:"comments"`
}

type Comment struct {
	Message       string             `json:"message" bson:"message"`
	CreatedBy     primitive.ObjectID `json:"-" bson:"created_by"`
	CreatedByUser Users              `json:"created_by" bson:"-"`
	CreatedAt     *time.Time         `json:"created_at" bson:"created_at"`
}

type Users struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
