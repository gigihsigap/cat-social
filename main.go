
package main

import (
	"cat-social/db"
	"cat-social/utils"
	"cat-social/routes"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
)

func main() {
	var err error
	var conn *pgx.Conn

	utils.LoadEnvVariables()

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	conn, err = db.ConnectToDatabase(dbURL)
	if err != nil {
		log.Fatal("DB connection failed!")
	}

	r := routes.SetupRouter(conn)

	r.Run(":8080")
}
