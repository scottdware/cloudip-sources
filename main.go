package main

import (
	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/aws", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "aws",
			"url":  "https://ip-ranges.amazonaws.com/ip-ranges.json",
		})
	})

	r.GET("/google", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "google",
			"url":  "https://www.gstatic.com/ipranges/goog.json",
		})
	})

	r.GET("/azure", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "azure",
			"url":  "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20220214.json",
		})
	})

	r.Run(":8080")
}
