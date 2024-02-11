package main

import (
	"net/http"
 	"github.com/gin-gonic/gin"
)


/* func main() {
	// Iniciada a Router utilizando as configurações Default do gin
	r *gin.Engine := gin.Default()
	// Definindo uma Rota
	r.GET(relativePath: "/ping", handlers... : func(c *gin.Context){
		c.JSON(code: 200, obj: gin.H{
			"message" : "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
 */

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }