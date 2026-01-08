package main

import (	
	"net/http"
	"googleCLS/config"
	"googleCLS/models"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
	config.ConnectDB() 

	r := gin.Default()
	// Set Gin mode to release for production
    // gin.SetMode(gin.ReleaseMode)

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

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

        courses, _ := models.GetCourses()
        slugMap := make(map[uint]string)
        for _, course := range courses {
            slugMap[course.ID] = course.Slug
        }

        for i := range assignments {
            assignments[i].CourseSlug = slugMap[assignments[i].CourseID]
        }

        c.JSON(http.StatusOK, assignments)
    })

    r.GET("/resources", func(c *gin.Context) {
        resources, err := models.GetResources()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return 
        }

        courses, _ := models.GetCourses()
        slugMap := make(map[uint]string)
        for _, course := range courses {
            slugMap[course.ID] = course.Slug
        }

        for i := range resources {
            resources[i].CourseSlug = slugMap[resources[i].CourseID]
        }

        c.JSON(http.StatusOK, resources)
    })

    r.GET("/course-users", func(c *gin.Context) {
        usersByCourse, err := models.GetUsersByCourse()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, usersByCourse)
    })
    

    r.Run(":9000")
}