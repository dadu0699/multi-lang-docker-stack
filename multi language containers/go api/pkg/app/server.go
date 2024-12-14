package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// NewServer creates a new server instance with the given dependencies (services)
func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

// Run starts the server on the given port and returns an error if something goes wrong during startup
func (server *Server) Run() error {
	// Run function that initializes the routes
	r := server.Routes()

	// Run the server through the router
	err := r.Run()
	if err != nil {
		log.Fatalf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
