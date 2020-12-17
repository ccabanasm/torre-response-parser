package jobs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"torre-response-parser/pkg/data"
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

// GetResults function unmarshalls a byte response and returns an array of Results
func GetResults(bytesJobs []byte) ([]Result, error) {
	var jobs Jobs

	err := json.Unmarshal(bytesJobs, &jobs)
	if err != nil {
		return nil, err
	}

	return jobs.Results, nil
}

// GetLastString is utilized to extract country for every job
func GetLastString(result []string) (string, error) {
	splitted := strings.Split(result[len(result)-1], ", ")
	return splitted[len(splitted)-1], nil
}

// OrganizationSave saves object to database
func OrganizationSave(organizations []Organization) (int, error) {
	stmt, err := data.Db.Prepare(`
		INSERT INTO organization (org_id, name, picture)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var orgID int
	err = stmt.QueryRow(organizations[0].Id, organizations[0].Name, organizations[0].Picture).Scan(&orgID)
	if err != nil {
		return 0, err
	}

	return orgID, nil
}

// CompensationSaves saves comp to database
func CompensationSave(compensation Compensation) (int, error) {
	stmt, err := data.Db.Prepare(`
		INSERT INTO compensation (code, currency, min_amount, max_amount, periodicity)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var compID int
	err = stmt.QueryRow(compensation.Data.Code, compensation.Data.Currency, compensation.Data.MinAmount, compensation.Data.MaxAmount, compensation.Data.Periodicity).Scan(&compID)
	if err != nil {
		return 0, err
	}

	return compID, nil
}

// MemberGetOrCreate returns member id plus error after creation
func MemberGetOrCreate(m Member) (int, error) {
	stmt, err := data.Db.Prepare(`
		INSERT INTO member (subject_id, username, prof_headline, picture, member, manager, poster, weight)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var memberID int
	err = stmt.QueryRow(m.SubjectId, m.Username, m.ProfessionalHeadline, m.Picture, m.Member, m.Manager, m.Poster, m.Weight).Scan(&memberID)
	if err != nil {
		return 0, err
	}

	return memberID, nil
}

func SkillSave(sk Skill, jobID int) error {
	stmt, err := data.Db.Prepare(`
		INSERT INTO skill_req (name, experience, job_id)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var skillID int
	err = stmt.QueryRow(sk.Name, sk.Experience, jobID).Scan(&skillID)
	if err != nil {
		return err
	}

	return nil
}

func QuestionSave(q Question, jobID int) error {
	stmt, err := data.Db.Prepare(`
		INSERT INTO question (id, text, date, job_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var questionID string
	err = stmt.QueryRow(q.Id, q.Text, q.Date, jobID).Scan(&questionID)
	if err != nil {
		return err
	}

	return nil
}

func JobSave(j Result, orgID int, compID int) (int, error) {
	country := ""
	if len(j.Locations) > 0 {
		c, err := GetLastString(j.Locations)
		if err != nil {
			return 0, err
		}
		country = c
	}
	stmt, err := data.Db.Prepare(`
		INSERT INTO job (job_id, objective, type, organization, locations, remote, external, deadline, status, compensation)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var jobID int
	err = stmt.QueryRow(j.Id, j.Objective, j.Type, orgID, country, j.Remote, j.External, j.Deadline, j.Status, compID).Scan(&jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

func SetJobToMember(jobID, memberID int) error {
	stmt, err := data.Db.Prepare(`
		INSERT INTO job_to_member (job_id, member_id)
		VALUES ($1, $2)
		RETURNING id;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(jobID, memberID).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func LoadToDB(results []Result) error {
	for i := 0; i < len(results); i++ {
		if i%100 == 0 {
			fmt.Println("Job N-", i)
		}
		orgID := 0
		if len(results[i].Organizations) > 0 {
			id, err := OrganizationSave(results[i].Organizations)
			if err != nil {
				return err
			}
			orgID = id
		}

		compID, err := CompensationSave(results[i].Compensation)
		if err != nil {
			return err
		}

		jobID, err := JobSave(results[i], orgID, compID)
		if err != nil {
			return err
		}

		for j := 0; j < len(results[i].Skills); j++ {
			err = SkillSave(results[i].Skills[j], jobID)
			if err != nil {
				return err
			}

		}

		for k := 0; k < len(results[i].Members); k++ {
			memberID, err := MemberGetOrCreate(results[i].Members[k])
			if err != nil {
				return err
			}
			err = SetJobToMember(jobID, memberID)
		}

		for l := 0; l < len(results[i].Questions); l++ {
			err = QuestionSave(results[i].Questions[l], jobID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
