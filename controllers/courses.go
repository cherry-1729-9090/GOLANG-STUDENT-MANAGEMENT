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

func GetCourses(c *gin.Context) {
	var courses []models.Course
	cursor, err := database.DB.Collection("courses").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching courses"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var course models.Course
		cursor.Decode(&course)
		courses = append(courses, course)
	}

	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("courses").InsertOne(context.TODO(), course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating course"})
		return
	}

	c.JSON(http.StatusOK, course)
}
