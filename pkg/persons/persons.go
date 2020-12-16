package persons

import (
	"encoding/json"
	"fmt"
	"strings"
	"torre-response-parser/pkg/data"
)

// Persons struct disaggregates results from json response
type Persons struct {
	Results []Result `json:"results"`
}

// Result struct unmarshalls a result object
type Result struct {
	Compensations Compensations `json:"compensations"`
	LocationName  string        `json:"locationName"`
	Name          string        `json:"name"`
	OpenTo        []string      `json:"openTo"`
	Picture       string        `json:"picture"`
	ProfHeadline  string        `json:"professionalHeadline"`
	Remoter       bool          `json:"remoter"`
	Skills        []Skill       `json:"skills"`
	SubjectId     string        `json:"subjectId"`
	Username      string        `json:"username"`
	Verified      bool          `json:"verified"`
	Weight        float64       `json:"weight"`
}

// Compensations struct stores compensations
type Compensations struct {
	Freelancer Compensation `json:"freelancer"`
	Employee   Compensation `json:"employee"`
	Intern     Compensation `json:"intern"`
}

// Compensation struct stores info for every Comp
type Compensation struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Periodicity string  `json:"periodicity"`
}

// Skills struct stores skill info
type Skill struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

// GetResults function unmarshalls a byte response and returns an array o Results
func GetResults(bytesJobs []byte) ([]Result, error) {
	var persons Persons

	err := json.Unmarshal(bytesJobs, &persons)
	if err != nil {
		return nil, err
	}

	return persons.Results, nil
}

func CompensationSave(c Compensation) (int, error) {
	stmt, err := data.Db.Prepare(`
		INSERT INTO person_compensation (amount, currency, periodicity)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var compID int
	err = stmt.QueryRow(c.Amount, c.Currency, c.Periodicity).Scan(&compID)
	if err != nil {
		return 0, err
	}

	return compID, nil
}

func CompensationsSave(c Compensations) (int, error) {
	freeID, err := CompensationSave(c.Freelancer)
	if err != nil {
		return 0, err
	}

	empID, err := CompensationSave(c.Employee)
	if err != nil {
		return 0, nil
	}

	intID, err := CompensationSave(c.Intern)
	if err != nil {
		return 0, nil
	}

	stmt, err := data.Db.Prepare(`
		INSERT INTO compensations (freelancer_id, employee_id, intern_id)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var compID int
	err = stmt.QueryRow(freeID, empID, intID).Scan(&compID)
	if err != nil {
		return 0, err
	}

	return compID, nil
}

// GetLastString is utilized to extract country for every person
func GetLastString(result string) (string, error) {
	splitted := strings.Split(result, ", ")
	return splitted[len(splitted)-1], nil
}

func PersonSave(p Result, compID int) (int, error) {
	country := ""
	if len(p.LocationName) > 0 {
		c, err := GetLastString(p.LocationName)
		if err != nil {
			return 0, err
		}
		country = c
	}

	stmt, err := data.Db.Prepare(`
		INSERT INTO person (compensations_id, location, name, picture, prof_headline, remoter, subject_id, username, verified, weight)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var perID int
	err = stmt.QueryRow(compID, country, p.Name, p.Picture, p.ProfHeadline, p.Remoter, p.SubjectId, p.Username, p.Verified, p.Weight).Scan(&perID)
	if err != nil {
		return 0, err
	}

	return perID, nil
}

func SkillSave(s Skill, perID int) error {
	stmt, err := data.Db.Prepare(`
		INSERT INTO skills (name, weight, person_id)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var skillID int
	err = stmt.QueryRow(s.Name, s.Weight, perID).Scan(&skillID)
	if err != nil {
		return err
	}

	return nil
}

func OpenToSave(o string, perID int) error {
	stmt, err := data.Db.Prepare(`
		INSERT INTO open_to_options (name, person_id)
		VALUES ($1, $2)
		RETURNING id;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var compID int
	err = stmt.QueryRow(o, perID).Scan(&compID)
	if err != nil {
		return err
	}

	return nil
}

func LoadToDB(results []Result) error {
	for i := 0; i < len(results); i++ {
		fmt.Println("Person N-", i)

		compID, err := CompensationsSave(results[i].Compensations)
		if err != nil {
			return err
		}

		perID, err := PersonSave(results[i], compID)
		if err != nil {
			return err
		}

		if len(results[i].Skills) > 0 {
			for j := 0; j < len(results[i].Skills); j++ {
				err = SkillSave(results[i].Skills[j], perID)
				if err != nil {
					return err
				}

			}
		}

		if len(results[i].OpenTo) > 0 {
			for k := 0; k < len(results[i].OpenTo); k++ {
				err := OpenToSave(results[i].OpenTo[k], perID)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}
