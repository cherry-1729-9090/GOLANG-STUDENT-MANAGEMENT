package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
    ID                primitive.ObjectID `bson:"_id,omitempty"`
    QuestionID        string             `bson:"question_id"`
    AssignmentID      primitive.ObjectID `bson:"assignment_id"` // Reference to the Assignment
    QuestionDescription string           `bson:"question_description"`
    Photos            []string           `bson:"photos,omitempty"` // Array of photo URLs
    Options           []string           `bson:"options"`         // Array of options
    CorrectOption     int                `bson:"correct_option"`  // Index of the correct option (0-3)
}
