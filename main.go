package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"torre-response-parser/pkg/data"
	"torre-response-parser/pkg/persons"
)

const (
	FILE_PATH = "/home/carlos/dev/torre/"
)

func main() {
	// database connection is initialized
	data.InitDb()

	// file string with path to json file
	file := FILE_PATH + "people.json"

	// jsonFile keeps json file opened
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// bytesJobs stores the whole json file in memory
	bytesJobs, _ := ioutil.ReadAll(jsonFile)

	// resultados var of type []Result stores json unmarshalled
	resultados, err := persons.GetResults(bytesJobs)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(resultados); i++ {
		fmt.Println("Person N-" + strconv.Itoa(i+1) + ": " + resultados[i].Username)
	}

}
