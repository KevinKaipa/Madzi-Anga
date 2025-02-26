package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files (CSS, JS, etc.)
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Admin dashboard route
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "Madzi Anga Admin Dashboard",
		})
	})

	// Run the server
	r.Run(":8080")
}
