package main

import (
	"golang-gin-api-rest/models"
	"golang-gin-api-rest/routes"
)

func main() {
	models.Students = []models.Student{
		{"Wesley", "12345678900", "123456780"},
		{"Astolfo", "12345678900", "123456780"},
	}
	routes.HandleRequests()
}
