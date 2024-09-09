package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Course struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    CourseID     string             `bson:"course_id"`
    CourseName   string             `bson:"course_name"`
    CourseFees   float64            `bson:"course_fees"`
    CourseDuration int              `bson:"course_duration"` // Duration in weeks or months
}