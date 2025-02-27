package service

import (
	"database/sql"

	controller "github.com/KentDavidButler/JobTracker/src/Controller"
	"github.com/gin-gonic/gin"
)

func GetJobPostings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.GetJobPostings(c, db)
	}
}

func GetJobPostingsByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.GetJobPostingsByID(c, db)
	}
}

func PostJobPosting(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.PostJobPosting(c, db)
	}
}

func Receiver() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.Receiver(c)
	}
}
