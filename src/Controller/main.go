package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/KentDavidButler/JobTracker/internal/receiver"
	model "github.com/KentDavidButler/JobTracker/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Job Postings

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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		fmt.Println("Error: ", err)
		return
	}

	newJobPosting.ID = uuid.NewString()

	model.SetJobPostings(newJobPosting, db)
	c.IndentedJSON(http.StatusCreated, newJobPosting)
}

type Receiver_Input struct {
	Url string `json:"url" binding:"required"` // required
}

func Receiver(c *gin.Context) {
	var url Receiver_Input

	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		fmt.Println("Error: ", err)
		return
	}

	receiver.Receiver(url.Url)
	c.IndentedJSON(http.StatusOK, gin.H{"url_received": url})

}

// Connections

func GetConnections(c *gin.Context, db *sql.DB) {
	jobs := model.GetConnections(db, 0)

	if len(jobs) > 0 {
		c.IndentedJSON(http.StatusOK, jobs)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "list is empty"})
}

func GetConnectionsByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	job := model.GetConnectionsByID(db, id)
	if job.ID != "" {
		c.IndentedJSON(http.StatusOK, job)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

func PostConnection(c *gin.Context, db *sql.DB) {
	var newConnection model.Connection

	if err := c.BindJSON(&newConnection); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		fmt.Println("Error: ", err)
		return
	}

	newConnection.ID = uuid.NewString()

	model.SetConnections(newConnection, db)
	c.IndentedJSON(http.StatusCreated, newConnection)
}

// Companies

func GetCompanies(c *gin.Context, db *sql.DB) {
	jobs := model.GetCompanies(db, 0)

	if len(jobs) > 0 {
		c.IndentedJSON(http.StatusOK, jobs)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "list is empty"})
}

func GetCompaniesByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	job := model.GetCompaniesByID(db, id)
	if job.ID != "" {
		c.IndentedJSON(http.StatusOK, job)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "company not found"})
}

func PostCompanies(c *gin.Context, db *sql.DB) {
	var newCompany model.Company

	if err := c.BindJSON(&newCompany); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		fmt.Println("Error: ", err)
		return
	}

	newCompany.ID = uuid.NewString()

	model.SetCompanies(newCompany, db)
	c.IndentedJSON(http.StatusCreated, newCompany)
}
