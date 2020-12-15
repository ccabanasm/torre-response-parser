package jobs

import (
	"encoding/json"
	"time"
)

// Jobs struct parses json response stored in file
type Jobs struct {
	Results []Result `json:"results"`
}

// Result struct disaggregates results from file
type Result struct {
	Id            string         `json:"id"`
	Objective     string         `json:"objective"`
	Type          string         `json:"type"`
	Organizations []Organization `json:"organizations"`
	Locations     []string       `json:"locations"`
	Remote        bool           `json:"remote"`
	External      bool           `json:"external"`
	Deadline      time.Time      `json:"deadline"`
	Status        string         `json:"status"`
	Compensation  Compensation   `json:"compensation"`
	Skills        []Skill        `json:"skills"`
	Members       []Member       `json:"members"`
	Questions     []Question     `json:"questions"`
}

// Organization struct stores an organization object
type Organization struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// Compensation struct stores a compensation object
type Compensation struct {
	Data    DataCompensation `json:"data"`
	Visible bool             `json:"visible"`
}

// DataCompensation struct disaggregates compensation data
type DataCompensation struct {
	Code        string  `json:"code"`
	Currency    string  `json:"currency"`
	MinAmount   float64 `json:"minAmount"`
	MaxAmount   float64 `json:"maxAmount"`
	Periodicity string  `json:"periodicity"`
}

// Skill struct stores a skill object
type Skill struct {
	Name       string `json:"name"`
	Experience string `json:"experience"`
}

// Member struct stores a member object
type Member struct {
	SubjectId            string  `json:"subjectId"`
	Name                 string  `json:"name"`
	Username             string  `json:"username"`
	ProfessionalHeadline string  `json:"professionalHeadline"`
	Picture              string  `json:"picture"`
	Member               bool    `json:"member"`
	Manager              bool    `json:"manager"`
	Poster               bool    `json:"poster"`
	Weight               float64 `json:"weight"`
}

// Question struct stores a question object
type Question struct {
	Id   string    `json:"id"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

func GetResults(bytesJobs []byte) ([]Result, error) {
	var jobs Jobs

	err := json.Unmarshal(bytesJobs, &jobs)
	if err != nil {
		return nil, err
	}

	return jobs.Results, nil
}
