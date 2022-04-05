package main

import "github.com/gin-gonic/gin"

/*
Author: Josiah Groux
This is a project that creates a simple Go webpage. It has a YouTube video imbedded in.
*/
func main() {
	r := gin.Default()
	// Loads in the templates in static
	r.LoadHTMLGlob("static/templates/*")
	// The following creates the first html page as default and I have set it as the duck.tmpl
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "duck.tmpl", gin.H{})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
