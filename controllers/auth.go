package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"lms-backend/database"
	"lms-backend/models"
	"lms-backend/utils"
)

func Register(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	student.Password = string(hashedPassword)

	student.ID = primitive.NewObjectID()
	_, err = database.DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating student"})
		return
	}

	token, _ := utils.GenerateJWT(student.ID.Hex(), "student")
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	err := database.DB.Collection("students").FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching student"})
		return
	}

	// Check if passwords match
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(student.ID.Hex(), "student")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
