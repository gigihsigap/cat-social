package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"cat-social/routes"
	"os"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup DB connection
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	err = setupDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	env := os.Getenv("ENC")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routes.SetupRouter(db)

	r.Run(":8080")
}

func setupDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)

	return err
}
