package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    TeacherID string             `bson:"teacher_id"`
}
