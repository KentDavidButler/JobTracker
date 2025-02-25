package service

import (
	controller "github.com/KentDavidButler/JobTracker/src/Controller"
	"github.com/gin-gonic/gin"
)

func GetJobPostings(c *gin.Context) {
	controller.GetJobPostings(c)
}

func GetJobPostingsByID(c *gin.Context) {
	controller.GetJobPostingsByID(c)
}

func PostJobPosting(c *gin.Context) {
	controller.PostJobPosting(c)
}
