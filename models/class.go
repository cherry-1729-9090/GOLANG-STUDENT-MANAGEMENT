package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    ClassID      string             `bson:"class_id"`
    BatchID      primitive.ObjectID `bson:"batch_id"`   // Reference to the Batch
    TeacherID    primitive.ObjectID `bson:"teacher_id"` // Reference to the Teacher
    ClassTopic   string             `bson:"class_topic"`
    Notes        string             `bson:"notes,omitempty"`
    ClassDuration int               `bson:"class_duration"` // Duration in minutes
    ClassTiming   string            `bson:"class_timing"` // ISO 8601 formatted time string
}
