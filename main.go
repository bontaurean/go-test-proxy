package main

import (
	"os"

	"github.com/bontaurean/go-test-proxy/handlers"
	"github.com/gin-gonic/gin"
)

const DEFAULT_PORT = "80"

func main() {
	server := setupServer(true)

	servicePort := os.Getenv("PORT")
	if servicePort == "" {
		servicePort = DEFAULT_PORT
	}

	server.Run(":" + servicePort)
}

func setupServer(testMode ...bool) *gin.Engine {
	e := gin.Default()

	r := e.Group("/v1/requests")
	{
		r.POST("/", handlers.HandleClientRequest)
		r.GET("/:requestId", handlers.HandleHistoryLookup)
	}

	if len(testMode) > 0 {
		e.GET("/test", handlers.HandleTestRequest)
	}

	return e
}
