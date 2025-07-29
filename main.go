package main

import (
	"io"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// /health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// /date endpoint
	r.GET("/date", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"date": time.Now().Format(time.RFC3339),
		})
	})

	// /print endpoint
	r.POST("/print", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
			return
		}
		// Assume incoming body is JSON, just return as is
		c.Data(http.StatusOK, "application/json", body)
	})

	// /shell endpoint
	r.POST("/shell", func(c *gin.Context) {
		out, err := exec.Command("uname", "-r").Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get kernel version"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"kernel": string(out),
		})
	})

	r.Run(":8080")
}
