package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/database"
	"github.com/hjunor/api-rest-golang.git/models"
	"github.com/hjunor/api-rest-golang.git/services"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json" + err.Error(),
		})
		return
	}
	user.Password = services.SHAR256Encoder(user.Password)
	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user" + err.Error(),
		})
		return
	}

	c.Status(204)
}
