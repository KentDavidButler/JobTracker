package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Manual Testing Curl Commands
// curl localhost:8080/jobpostings
// curl -d '{"id":"123", "companyName":"some_Company2", "positionLink":"www.some_company2.com"}' -H "Content-Type: application/json" -X POST http://localhost:8080/jobpostings

type JobPosting struct {
	ID                        string  `json:"id,omitempty"`
	CompanyName               string  `json:"companyName" binding:"required"` // required
	ReferralName              string  `json:"referralName,omitempty"`
	ReferralNotes             float64 `json:"referralNotes,omitempty"`
	ApplicationSubmissionDate string  `json:"applicationSubmissionDate,omitempty"`
	PositionLink              string  `json:"positionLink" binding:"required"` // required
	GoogleDocLink             string  `json:"docLink,omitempty"`
	Interview                 bool    `json:"interview,omitempty"`
	InterviewDate             string  `json:"interviewDate,omitempty"`
	Denial                    bool    `json:"denial,omitempty"`
	AdditionalInfo            string  `json:"additionalInfo,omitempty"`
}

var jobPostings = []JobPosting{
	{ID: "39cb5563-f85a-43d2-a815-51ced1138b9f", CompanyName: "someCompany", PositionLink: "www.somecompany.com"},
}

func getJobPostings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobPostings)
}

func getJobPostingsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range jobPostings {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postJobPosting(c *gin.Context) {
	var newJobPosting JobPosting

	if err := c.BindJSON(&newJobPosting); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if newJobPosting.CompanyName == "" || newJobPosting.PositionLink == "" {
		gin.ErrorLogger()
		return
	}

	newJobPosting.ID = uuid.NewString()

	jobPostings = append(jobPostings, newJobPosting)
	c.IndentedJSON(http.StatusCreated, newJobPosting)
}

func main() {
	router := gin.Default()
	router.GET("/jobpostings", getJobPostings)
	router.GET("/jobpostings/:id", getJobPostingsByID)
	router.POST("/jobpostings", postJobPosting)

	router.Run("localhost:8080")
}
