package handlers

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task struct to represent a task object
type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	UserID      string             `json:"user_id" bson:"user_id"`
}
