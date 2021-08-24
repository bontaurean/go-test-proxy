package handlers

import (
	"net/http"

	"github.com/bontaurean/go-test-proxy/fetcher"
	"github.com/bontaurean/go-test-proxy/models"
	"github.com/bontaurean/go-test-proxy/storage"
	"github.com/gin-gonic/gin"
)

func HandleClientRequest(c *gin.Context) {
	var req models.ProxyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	r, err := fetcher.Fetch(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

func HandleHistoryLookup(c *gin.Context) {
	requestId := c.Param("requestId")

	presp, err := storage.History.Get(requestId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, presp)
}
