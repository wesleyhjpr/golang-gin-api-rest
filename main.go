package main

import (
	"golang-gin-api-rest/database"
	"golang-gin-api-rest/models"
	"golang-gin-api-rest/routes"
)

func main() {
	database.ConnectToDatabase()

	models.Students = []models.Student{
		{Name: "Wesley", CPF: "12345678900", RG: "123456780"},
		{Name: "Astolfo", CPF: "12345678900", RG: "123456780"},
	}
	routes.HandleRequests()
}
