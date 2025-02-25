package main

import (
	service "github.com/KentDavidButler/JobTracker/src/Service"
	"github.com/gin-gonic/gin"
)

// Manual Testing Curl Commands
// curl localhost:8080/jobpostings
// curl -d '{"id":"123", "companyName":"some_Company2", "positionLink":"www.some_company2.com"}' -H "Content-Type: application/json" -X POST http://localhost:8080/jobpostings

func main() {
	router := gin.Default()
	router.GET("/jobpostings", service.GetJobPostings)
	router.GET("/jobpostings/:id", service.GetJobPostingsByID)
	router.POST("/jobpostings", service.PostJobPosting)

	router.Run("localhost:8080")
}
