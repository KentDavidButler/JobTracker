package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	service "github.com/KentDavidButler/JobTracker/src/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// TODO: Standardize output for errors and success
func main() {
	fmt.Println("Starting Service")
	// Pull ENV VARS
	pg_user, pg_pass, db_name := pullEvnVarsOrExit()

	db_conn_string := fmt.Sprintf("host=127.0.0.1 port=5432 user=%s password =%s dbname=%s sslmode=disable", pg_user, pg_pass, db_name)
	db, err := sql.Open("postgres", db_conn_string)
	if err != nil {
		panic("DB Open: " + err.Error())
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/jobpostings", service.GetJobPostings(db))
	router.GET("/jobpostings/:id", service.GetJobPostingsByID(db))
	router.POST("/jobpostings", service.PostJobPosting(db))
	// TODO: Make a put endpoint to update a job posting

	// Pull Data from Job Posting URL
	router.POST("/jobpostings/receiver", service.Receiver())

	// Add Connections/People/Referrers to DB
	router.GET("/connections", service.GetConnections(db))
	router.GET("/connections/:id", service.GetConnectionsByID(db))
	router.POST("/connetions", service.PostConnection(db))

	// Add Connections/People/Referrers to DB
	router.GET("/companies", service.GetCompanies(db))
	router.GET("/companies/:id", service.GetCompaniesByID(db))
	router.POST("/companies", service.PostCompanies(db))

	router.Run("localhost:8080")
}

func pullEvnVarsOrExit() (string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	pg_user, present := os.LookupEnv("POSTGRES_USERNAME")
	if present {
		fmt.Println("pg user name found")
	} else {
		panic("pg user name not found")

	}
	pg_pass, present := os.LookupEnv("POSTGRES_PASSWORD")
	if present {
		fmt.Println("pg user password found")
	} else {
		panic("pg user password not found")
	}
	db_name, present := os.LookupEnv("POSTGRES_DBNAME")
	if present {
		fmt.Println("pg db name password found")
	} else {
		panic("pg db name password not found")
	}

	return pg_user, pg_pass, db_name
}
