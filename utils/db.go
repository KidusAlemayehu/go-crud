package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() {
	var err error
	godotenv.Load()
	databaseName := os.Getenv("DB_NAME")
	databaseUsername := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseHost := os.Getenv("DB_HOST")
	databasePort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", databaseUsername, databasePassword, databaseName, databaseHost, databasePort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected to database")
}
