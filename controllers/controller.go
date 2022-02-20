package controllers

import (
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says:": "what's up " + name,
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
