package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type JobPosting struct {
	ID                        string `json:"id,omitempty"`
	CompanyName               string `json:"companyName" binding:"required"` // required
	ReferralName              string `json:"referralName,omitempty"`
	ReferralNotes             string `json:"referralNotes,omitempty"`
	ApplicationSubmissionDate string `json:"applicationSubmissionDate,omitempty"`
	PositionLink              string `json:"positionLink" binding:"required"` // required
	GoogleDocLink             string `json:"docLink,omitempty"`
	Interview                 bool   `json:"interview,omitempty"`
	InterviewDate             string `json:"interviewDate,omitempty"`
	Denial                    bool   `json:"denial,omitempty"`
	AdditionalInfo            string `json:"additionalInfo,omitempty"`
}

var JobPostings = []JobPosting{
	{ID: "39cb5563-f85a-43d2-a815-51ced1138b9f", CompanyName: "someCompany", PositionLink: "www.somecompany.com"},
}

func GetJobPostings(db *sql.DB) []JobPosting {
	return JobPostings
}

func GetJobPostingsByID(db *sql.DB, id string) JobPosting {
	var job JobPosting

	// this only works because all DB fields map to all JobPosting struct fields
	// if for any reason they don't match in the future, we'll have to request each column individually
	// this also works because fields are not able to be null in the DB and have default values.
	err := db.QueryRow("SELECT * FROM job_postings WHERE id = $1", id).Scan(&job.ID, &job.CompanyName,
		&job.ReferralName, &job.ReferralNotes, &job.ApplicationSubmissionDate, &job.PositionLink,
		&job.GoogleDocLink, &job.Interview, &job.InterviewDate, &job.Denial, &job.AdditionalInfo)
	if err != nil {
		log.Fatal(err)
	}

	return job
}

func SetJobPostings(posting JobPosting, db *sql.DB) {
	JobPostings = append(JobPostings, posting)
}
