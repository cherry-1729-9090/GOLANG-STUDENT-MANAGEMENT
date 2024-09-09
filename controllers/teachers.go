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

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	cursor, err := database.DB.Collection("teachers").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching teachers"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var teacher models.Teacher
		cursor.Decode(&teacher)
		teachers = append(teachers, teacher)
	}

	c.JSON(http.StatusOK, teachers)
}

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("teachers").InsertOne(context.TODO(), teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating teacher"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}
