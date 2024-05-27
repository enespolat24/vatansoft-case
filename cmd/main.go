package main

import (
	"log"
	"vatansoft-case/internal/database"
	routes "vatansoft-case/internal/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	e := echo.New()

	dbConn := database.NewDbConnection()
	defer dbConn.CloseDbConnection()

	if err := dbConn.AutoMigrate(); err != nil {
		e.Logger.Fatalf("Failed to migrate database: %v", err)
	}

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
