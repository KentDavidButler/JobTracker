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

func GetJobPostings(db *sql.DB, offset int16) []JobPosting {
	var postings []JobPosting

	rows, err := db.Query(`SELECT * FROM job_postings ORDER BY company_name LIMIT 
		100 OFFSET $1`, offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var job JobPosting
		err := rows.Scan(&job.ID, &job.CompanyName,
			&job.ReferralName, &job.ReferralNotes, &job.ApplicationSubmissionDate, &job.PositionLink,
			&job.GoogleDocLink, &job.Interview, &job.InterviewDate, &job.Denial, &job.AdditionalInfo)
		if err != nil {
			log.Fatal(err)
		}
		postings = append(postings, job)
	}

	return postings
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

func SetJobPostings(job JobPosting, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO job_postings(id, company_name, referral_name, referral_notes,
		application_submit_date, position_link, goog_doc_link, interview, interview_date, denial,
		 additional_info) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(&job.ID, &job.CompanyName, &job.ReferralName, &job.ReferralNotes,
		&job.ApplicationSubmissionDate, &job.PositionLink, &job.GoogleDocLink, &job.Interview,
		&job.InterviewDate, &job.Denial, &job.AdditionalInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB Response:  %d\n", res)
}
