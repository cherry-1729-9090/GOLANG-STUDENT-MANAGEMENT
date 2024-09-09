package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    SubjectID  string             `bson:"subject_id"`
    SubjectName string            `bson:"subject_name"`
}
