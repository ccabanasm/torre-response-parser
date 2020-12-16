package persons

import "encoding/json"

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
