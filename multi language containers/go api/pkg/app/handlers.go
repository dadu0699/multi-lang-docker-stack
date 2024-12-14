package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Root() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		response := gin.H{
			"message": "Welcome to the Go Gin API!",
		}
		c.JSON(http.StatusOK, response)
	}
}
