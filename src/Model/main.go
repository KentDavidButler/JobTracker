package model

type JobPosting struct {
	ID                        string  `json:"id,omitempty"`
	CompanyName               string  `json:"companyName" binding:"required"` // required
	ReferralName              string  `json:"referralName,omitempty"`
	ReferralNotes             float64 `json:"referralNotes,omitempty"`
	ApplicationSubmissionDate string  `json:"applicationSubmissionDate,omitempty"`
	PositionLink              string  `json:"positionLink" binding:"required"` // required
	GoogleDocLink             string  `json:"docLink,omitempty"`
	Interview                 bool    `json:"interview,omitempty"`
	InterviewDate             string  `json:"interviewDate,omitempty"`
	Denial                    bool    `json:"denial,omitempty"`
	AdditionalInfo            string  `json:"additionalInfo,omitempty"`
}

var JobPostings = []JobPosting{
	{ID: "39cb5563-f85a-43d2-a815-51ced1138b9f", CompanyName: "someCompany", PositionLink: "www.somecompany.com"},
}

func GetJobPostings() []JobPosting {
	return JobPostings
}

func SetJobPostings(posting JobPosting) {
	JobPostings = append(JobPostings, posting)
}
