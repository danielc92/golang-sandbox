package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func runGinApi() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		res := []Tag{
			{Name: "Arts", Description: "Tag for arts and crafts."},
		}
		c.JSON(http.StatusOK, res)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}