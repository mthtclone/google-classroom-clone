package main

import (	
	"net/http"
	"googleCLS/config"
	"googleCLS/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() 

	r := gin.Default()
	// Set Gin mode to release for production
    // gin.SetMode(gin.ReleaseMode)

	r.GET("/users", func(c *gin.Context) {
		users, err := models.GetUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
		}
		c.JSON(http.StatusOK, users)
	})

	r.GET("/courses", func(c *gin.Context) {
		courses, err := models.GetCourses()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, courses)
	})

	r.GET("/assignments", func(c *gin.Context) {
        assignments, err := models.GetAssignments()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, assignments)
    })

    r.GET("/resources", func(c *gin.Context) {
        resources, err := models.GetResources()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, resources)
    })

    r.Run(":8080")
}