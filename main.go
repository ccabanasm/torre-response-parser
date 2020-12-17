package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"torre-response-parser/pkg/data"
	"torre-response-parser/pkg/jobs"
)

const (
	FILE_PATH = "/home/carlos/dev/torre/"
)

func main() {
	// database connection is initialized
	data.InitDb()

	for k := 0; k < 4; k++ {
		// file string with path to json file
		file := FILE_PATH + "jobs-0" + strconv.Itoa(k) + ".json"

		fmt.Println("File: ", file)

		// jsonFile keeps json file opened
		jsonFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		// bytesJobs stores the whole json file in memory
		bytesJobs, _ := ioutil.ReadAll(jsonFile)

		// resultados var of type []Result stores json unmarshalled
		resultados, err := jobs.GetResults(bytesJobs)
		if err != nil {
			fmt.Println(err)
		}

		err = jobs.LoadToDB(resultados)
		if err != nil {
			fmt.Println(err)
		}

		jsonFile.Close()

	}

}
