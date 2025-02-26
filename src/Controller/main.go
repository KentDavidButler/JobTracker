package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	model "github.com/KentDavidButler/JobTracker/src/Model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetJobPostings(c *gin.Context, db *sql.DB) {
	jobs := model.GetJobPostings(db, 0)

	if len(jobs) > 0 {
		c.IndentedJSON(http.StatusOK, jobs)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "list is empty"})
}

func GetJobPostingsByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	job := model.GetJobPostingsByID(db, id)
	if job.ID != "" {
		c.IndentedJSON(http.StatusOK, job)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

func PostJobPosting(c *gin.Context, db *sql.DB) {
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

	model.SetJobPostings(newJobPosting, db)
	c.IndentedJSON(http.StatusCreated, newJobPosting)
}
