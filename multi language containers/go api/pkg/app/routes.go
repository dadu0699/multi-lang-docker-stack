package app

import "github.com/gin-gonic/gin"

// Routes sets up the routes for the application.
func (server *Server) Routes() *gin.Engine {
	router := server.router

	router.GET("/", server.Root())

	return router
}
