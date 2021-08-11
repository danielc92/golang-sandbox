package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Long float32 
	Lat float32 
	Name  string
}

type Student struct {
	
		FirstName string 
		LastName string 
	
}

type StudentApiResponse struct {
	Data []Student `json:"data"`
	Total int `json:"total"`
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

func LoadStudents() []Student {
	students := []Student{
		{"DANIEL","CORCORAN"},
		{"DANIEL","LAY"},
		{"DANIEL","LUKAS"},
		{"DANIEL","JULES"},
		{"SMITH","FIELDING"},
		{"SMITH","DOWNEY"},
		{"SMITH","JAN"},
		{"ANNA","RAY"},
		{"BELLE","RUFUS"},
		{"ROBERT","SMITH"},
	}
	return students
}
func runGinApi() {

	r := gin.Default()

	// r.Use(DanielsCustomMiddleware())
	// r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {

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
	

	r.GET("/students", func(c *gin.Context) {

		students := LoadStudents()
	
		firstName := c.Query("firstName")
	
		var results []Student
	
		for _, s := range students {
			if (s.FirstName == strings.ToUpper(firstName)) {
				results = append(results, s)
			}
		}
	
		 apiResponse := StudentApiResponse{
				Data: results,
				Total: len(results),
			}
	
	
		 
		c.JSON(http.StatusOK, &apiResponse)

	})

	r.PUT("/students", func(c *gin.Context) {

			oldStudent := Student{
				"Richard",
				"Smith",
			}

			var newStudent Student

			err := c.BindJSON(&newStudent)

			if (err != nil) {
				c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect body formatt!"})
				c.Abort()
			}

			if len(newStudent.FirstName) > 0 {
				oldStudent.FirstName = newStudent.FirstName
			}
			
			if len(newStudent.LastName) > 0 {
				oldStudent.LastName = newStudent.LastName
			}

			c.JSON(http.StatusOK, oldStudent)

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}