package controller

import (
	"fmt"
	"net/http"

	model "github.com/KentDavidButler/JobTracker/src/Model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetJobPostings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.JobPostings)
}

func GetJobPostingsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range model.GetJobPostings() {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostJobPosting(c *gin.Context) {
	var newJobPosting model.JobPosting

	if err := c.BindJSON(&newJobPosting); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if newJobPosting.CompanyName == "" || newJobPosting.PositionLink == "" {
		gin.ErrorLogger()
		return
	}

	newJobPosting.ID = uuid.NewString()

	model.SetJobPostings(newJobPosting)
	c.IndentedJSON(http.StatusCreated, newJobPosting)
}
