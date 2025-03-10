package service

import (
	"database/sql"

	controller "github.com/KentDavidButler/JobTracker/src/controller"
	"github.com/gin-gonic/gin"
)

// Job Postings

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

// Connections

func GetConnections(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.GetConnections(c, db)
	}
}

func GetConnectionsByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.GetConnectionsByID(c, db)
	}
}

func PostConnection(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.PostConnection(c, db)
	}
}
