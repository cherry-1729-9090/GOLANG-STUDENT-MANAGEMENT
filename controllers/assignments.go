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

func GetAssignments(c *gin.Context) {
	var assignments []models.Assignment
	cursor, err := database.DB.Collection("assignments").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching assignments"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var assignment models.Assignment
		cursor.Decode(&assignment)
		assignments = append(assignments, assignment)
	}

	c.JSON(http.StatusOK, assignments)
}

func CreateAssignment(c *gin.Context) {
	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("assignments").InsertOne(context.TODO(), assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating assignment"})
		return
	}

	c.JSON(http.StatusOK, assignment)
}
