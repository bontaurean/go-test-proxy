package main

import (
	"os"

	"github.com/bontaurean/go-test-proxy/handlers"
	"github.com/gin-gonic/gin"
)

const DEFAULT_PORT = "80"

func main() {
	server := initServer()

	servicePort := os.Getenv("PORT")
	if servicePort == "" {
		servicePort = DEFAULT_PORT
	}

	server.Run(":" + servicePort)
}

func initServer() *gin.Engine {
	e := gin.Default()

	r := e.Group("/v1/requests")
	{
		r.POST("/", handlers.HandleClientRequest)
		r.GET("/:requestId", handlers.HandleHistoryLookup)
	}

	return e
}
