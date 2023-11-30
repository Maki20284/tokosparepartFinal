package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", nil)
	})
	router.GET("/honda", func(c *gin.Context) {
		c.HTML(http.StatusOK, "honda.html", nil)
	})
	router.GET("/toyota", func(c *gin.Context) {
		c.HTML(http.StatusOK, "toyota.html", nil)
	})
	router.GET("/daihatsu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "daihatsu.html", nil)
	})

	router.Run(":8080")
}
