package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTestRequest(c *gin.Context) {
	// requestId := c.Query("cl")
	// c.Request.Response.Header.Del("")
	c.String(http.StatusOK, "999")
	// c.Writer.Header().Del("Content-Length")
	c.Header("Content-Length", "-1")
	fmt.Printf("%#v", c.Writer.Header())
}
