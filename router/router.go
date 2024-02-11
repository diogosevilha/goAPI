package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	router.Run(addr... : ":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }