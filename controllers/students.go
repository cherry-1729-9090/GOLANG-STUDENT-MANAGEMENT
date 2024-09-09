package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lms-backend/database"
	"lms-backend/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	cursor, err := database.DB.Collection("students").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching students"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var student models.Student
		cursor.Decode(&student)
		students = append(students, student)
	}

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating student"})
		return
	}

	c.JSON(http.StatusOK, student)
}
