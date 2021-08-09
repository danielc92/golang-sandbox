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

func runGinApi() {
	r := gin.Default()
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