package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Shaufi Imanulhaq",
			"bio":  "A Software engineer & content creator",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "Belajar Golang bareng Shaufi Imanulhaq",
		})
	})

	router.Run(":8888")
}
