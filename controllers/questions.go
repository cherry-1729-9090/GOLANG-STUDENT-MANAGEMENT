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

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	cursor, err := database.DB.Collection("questions").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching questions"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var question models.Question
		cursor.Decode(&question)
		questions = append(questions, question)
	}

	c.JSON(http.StatusOK, questions)
}

func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("questions").InsertOne(context.TODO(), question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating question"})
		return
	}

	c.JSON(http.StatusOK, question)
}
