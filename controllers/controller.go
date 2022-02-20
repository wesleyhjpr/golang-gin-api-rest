package controllers

import (
	"golang-gin-api-rest/models"

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
