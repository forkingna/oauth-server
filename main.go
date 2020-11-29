package main

import (
	"github.com/gin-gonic/gin"
)

func TestPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"code": 0,
	})
}

func main() {
	r := gin.Default()
	r.GET("ping", TestPing)
	r.Run("0.0.0.0:9301")
}
