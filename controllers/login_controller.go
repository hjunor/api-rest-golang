package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/database"
	"github.com/hjunor/api-rest-golang.git/models"
	"github.com/hjunor/api-rest-golang.git/services"
)

func Login(c *gin.Context) {

	db := database.GetDatabase()

	var p models.Login
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{"cannot bind JSON": err.Error()})
		return
	}

	var user models.User

	dbError := db.Where("email = ?", p.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	if user.Password != services.SHAR256Encoder(p.Password) {
		c.JSON(401, gin.H{"error": "Wrong password, ivalid credentias"})
		return
	}

	token, err := services.NewJwtServices().GenerateToken(user.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
