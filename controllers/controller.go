package controllers

import (
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says:": "what's up " + name,
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
