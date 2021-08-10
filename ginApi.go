package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Long float32 
	Lat float32 
	Name  string
}

/*
This middleware creates a flag (authenticated)
If authenticated evaluate to true, then execute the next handler in the chain
Otherwise send status unauthorized (401) code and prevent further chaining
*/

func DanielsCustomMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		authenticated := true
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{})
			// Immediate stop, dont call next handler
			c.Abort()
		} 
		c.Next()
	}
}

func runGinApi() {
	r := gin.Default()

	r.Use(DanielsCustomMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		// res := []Tag{
		// 	{Name: "Arts", Description: "Tag for arts and crafts."},
		// }
		res := gin.H{
			"x":1,
		}
	
		c.JSON(http.StatusOK, res)
	})

	r.GET("/geo", func(c *gin.Context) {
		res := []Country{
			{ 15.552727, 48.516388, "Yemen"},
			{ -19.015438, 29.154857, "Zimbabwe"},
			{ -13.133897, 27.849332, "Zambia"},
		}
		
		c.JSON(http.StatusOK, res)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}