package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Assignment struct {
    ID      primitive.ObjectID `bson:"_id,omitempty"`
    ClassID primitive.ObjectID `bson:"class_id"` // Reference to the Class
}
