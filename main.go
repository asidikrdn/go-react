package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	err := godotenv.Load("config.env")
	if err != nil {
		log.Println("Error loading environment variables file, the apps will read global environtment variabels on this system : ", err)
	}

	// Create a channel for graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// client
	clientService := ClientService{
		Path: "./client",
		Port: os.Getenv("FE_PORT"),
	}

	// server
	serverService := ServerService{
		Path: "./server",
		Port: os.Getenv("BE_PORT"),
	}

	if os.Getenv(gin.EnvGinMode) == gin.ReleaseMode {
		// production mode
		go clientService.RunProduction()
		go serverService.RunServer(gin.ReleaseMode)
	} else {
		// development mode
		go clientService.RunDevelopment()
		go serverService.RunServer(gin.DebugMode)
	}

	// Wait for a signal to gracefully shutdown
	<-shutdownChan
	log.Println("Received shutdown signal. Performing graceful shutdown...")

	// create context with timeout ** second
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the clientService
	clientService.Shutdown(ctx)

	// Gracefully shutdown the serverService
	serverService.Shutdown(ctx)

	// catching ctx.Done()
	<-ctx.Done()
	log.Println("Server and client gracefully shut down.")
}
