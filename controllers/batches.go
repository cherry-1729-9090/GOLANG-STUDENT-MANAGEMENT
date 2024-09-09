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

func GetBatches(c *gin.Context) {
	var batches []models.Batch
	cursor, err := database.DB.Collection("batches").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching batches"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var batch models.Batch
		cursor.Decode(&batch)
		batches = append(batches, batch)
	}

	c.JSON(http.StatusOK, batches)
}

func CreateBatch(c *gin.Context) {
	var batch models.Batch
	if err := c.ShouldBindJSON(&batch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	batch.ID = primitive.NewObjectID()
	_, err := database.DB.Collection("batches").InsertOne(context.TODO(), batch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating batch"})
		return
	}

	c.JSON(http.StatusOK, batch)
}
