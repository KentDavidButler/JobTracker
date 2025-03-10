package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type JobPosting struct {
	ID                        string       `json:"id,omitempty"`
	CompanyName               string       `json:"companyName" binding:"required"` // required
	ReferralName              string       `json:"referralName,omitempty"`
	ReferralNotes             string       `json:"referralNotes,omitempty"`
	Connections               []Connection `json:"connections,omitempty`
	ApplicationSubmissionDate string       `json:"applicationSubmissionDate,omitempty"`
	PositionLink              string       `json:"positionLink" binding:"required"` // required
	GoogleDocLink             string       `json:"docLink,omitempty"`
	Interview                 bool         `json:"interview,omitempty"`
	InterviewDate             string       `json:"interviewDate,omitempty"`
	Denial                    bool         `json:"denial,omitempty"`
	AdditionalInfo            string       `json:"additionalInfo,omitempty"`
}

type Connection struct {
	ID          string    `json:"id,omitempty"`
	FirstName   string    `json:"firstname" binding:"required"` // required
	LastName    string    `json:"lastname" binding:"required"`  // required
	Companies   []Company `json:"companies,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Email       string    `json:"email,omitempty"`
	LinkedInUrl string    `json:"linkedin_url,omitempty"`
}

type Company struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name" binding:"required"` // required
	Phone       string       `json:"phone,omitempty"`
	Connections []Connection `json:"connections,omitempty"`
	LinkedInUrl string       `json:"linkedin_url,omitempty"`
}

// Job Postings

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

// Connections

func GetConnections(db *sql.DB, offset int16) []Connection {
	var connections []Connection

	rows, err := db.Query(`SELECT * FROM connections ORDER BY name LIMIT 
		100 OFFSET $1`, offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var conn Connection
		err := rows.Scan(&conn.ID, &conn.FirstName, &conn.LastName,
			&conn.Companies, &conn.Phone, &conn.Email, &conn.LinkedInUrl)
		if err != nil {
			log.Fatal(err)
		}
		connections = append(connections, conn)
	}

	return connections
}

func GetConnectionsByID(db *sql.DB, id string) Connection {
	var conn Connection

	err := db.QueryRow("SELECT * FROM connections WHERE id = $1", id).Scan(&conn.ID,
		&conn.FirstName, &conn.LastName,
		&conn.Companies, &conn.Phone,
		&conn.Email, &conn.LinkedInUrl)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func SetConnections(conn Connection, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO connections(id, first_name, last_name, companies,
		phone, email, linked_in_url) VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(&conn.ID, &conn.FirstName, &conn.LastName,
		&conn.Companies, &conn.Phone, &conn.Email, &conn.LinkedInUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB Response:  %d\n", res)
}

// Companies

func GetCompanies(db *sql.DB, offset int16) []Company {
	var companies []Company

	rows, err := db.Query(`SELECT * FROM companies ORDER BY name LIMIT 
		100 OFFSET $1`, offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var comps Company
		err := rows.Scan(&comps.ID, &comps.Name, &comps.Phone,
			&comps.Connections, &comps.LinkedInUrl)
		if err != nil {
			log.Fatal(err)
		}
		companies = append(companies, comps)
	}

	return companies
}

func GetCompaniesByID(db *sql.DB, id string) Company {
	var comp Company

	err := db.QueryRow("SELECT * FROM companies WHERE id = $1", id).Scan(&comp.ID,
		&comp.Name, &comp.Phone, &comp.Connections, &comp.LinkedInUrl)
	if err != nil {
		log.Fatal(err)
	}

	return comp
}

func SetCompanies(comp Company, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO companies(id, name, phone, connections,
		linked_in_url) VALUES($1, $2, $3, $4, $5)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(&comp.ID, &comp.Name, &comp.Phone,
		&comp.Connections, &comp.LinkedInUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB Response:  %d\n", res)
}
