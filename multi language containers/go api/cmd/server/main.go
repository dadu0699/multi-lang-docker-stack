package main

import (
	"log"
	"os"

	"multicontainers/goapi/pkg/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func run() error {
	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	server := app.NewServer(router)
	if err := server.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	if err := run(); err != nil {
		log.Fatalf("This is the startup error: %v", err)
		os.Exit(1)
	}
}
