package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/database"
	"github.com/hjunor/api-rest-golang.git/models"
)

func ShowBook(c *gin.Context) {
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

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json" + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find list books" + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json" + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func DeleteBook(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger",
		})
		return
	}

	var book models.Book

	err = db.Delete(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": " cannot find book:  " + err.Error(),
		})

		return
	}

	c.JSON(200, book)

}
