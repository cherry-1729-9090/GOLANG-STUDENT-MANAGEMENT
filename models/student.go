package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    StudentID string             `bson:"student_id"`
    Email     string             `bson:"email"`
    Password  string             `bson:"password"` // Ensure to hash this
    BatchID   primitive.ObjectID `bson:"batch_id"` // Reference to the Batch
}
