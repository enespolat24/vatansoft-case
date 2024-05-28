package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vatansoft-case/internal/database"
	"vatansoft-case/internal/handler"
	"vatansoft-case/internal/repository"
	routes "vatansoft-case/internal/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	planRepo := repository.NewPlanRepository(dbConn.GetDbInstance())
	userHandler := handler.NewUserHandler(userRepo)
	authHandler := handler.NewAuthHandler(userRepo)
	planHandler := handler.NewPlanHandler(planRepo)

	routes.InitUserRoutes(e, userHandler)
	routes.InitAuthRoutes(e, authHandler)
	routes.InitPlanRoutes(e, planHandler)

	//prometheus
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	e.Logger.Print("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Print("Server shutdown completed")
}
