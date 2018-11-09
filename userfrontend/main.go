// Package of main frontend
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/img", "./img")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fe1.tmpl", gin.H{
			"title":   "RedHat Istio Demo",
			"version": "1.0.17",
		})
	})

	router.GET("/dbcount", func(c *gin.Context) {

	})
	router.Run(":8080")
}
