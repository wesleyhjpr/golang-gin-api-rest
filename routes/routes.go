package routes

import (
	"golang-gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.Run()
}
