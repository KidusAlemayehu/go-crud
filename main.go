package main

import (
	"crud/go-crud/models"
	"crud/go-crud/utils"
	"crud/go-crud/views"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var user models.User
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	utils.DBConn()
	utils.DB.AutoMigrate(&user)
}
func main() {
	router := gin.Default()

	router.GET("/", views.UsersList)
	router.GET("/user/:id", views.UserList)
	router.POST("/register", views.UserCreate)
	router.POST("/login", views.UserLogin)
	router.DELETE("/user/:id", views.UserRemove)

	router.Run("localhost:8080")
}
