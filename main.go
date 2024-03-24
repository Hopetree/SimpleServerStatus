// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 定义一个简单的API路由
	router.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Golang backend!"})
	})

	_ = router.Run(":8011")
}
