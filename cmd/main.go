package main

import (
	"log"
	"vatansoft-case/internal/database"
	"vatansoft-case/internal/handler"
	"vatansoft-case/internal/repository"
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

	userRepo := repository.NewUserRepository(dbConn.GetDbInstance())
	userHandler := handler.NewUserHandler(userRepo)
	authHandler := handler.NewAuthHandler(userRepo)

	routes.InitUserRoutes(e, userHandler)
	routes.InitAuthRoutes(e, authHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
