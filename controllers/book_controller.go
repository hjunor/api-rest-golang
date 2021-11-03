package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/database"
	"github.com/hjunor/api-rest-golang.git/models"
)

func ShowBooks(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger",
		})
		return
	}

	db := database.GetDatabase()
	var book models.Book

	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": " cannot find book:  " + err.Error(),
		})

		return
	}

	c.JSON(200, book)

}
