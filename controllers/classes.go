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

func GetClasses(c *gin.Context) {
	var classes []models.Class
	cursor, err := database.DB.Collection("classes").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching classes"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var class models.Class
		cursor.Decode(&class)
		classes = append(classes, class)
	}

	c.JSON(http.StatusOK, classes)
}

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("classes").InsertOne(context.TODO(), class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating class"})
		return
	}

	c.JSON(http.StatusOK, class)
}
