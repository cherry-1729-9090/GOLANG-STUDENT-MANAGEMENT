package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Batch struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    BatchID  string             `bson:"batch_id"`
    BatchYear int               `bson:"batch_year"`
    CourseID primitive.ObjectID `bson:"course_id"` // Reference to the Course
}
