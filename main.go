package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	file := FILE_PATH + "people-00.json"

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

	err = persons.LoadToDB(resultados)
	if err != nil {
		fmt.Println(err)
	}
}
