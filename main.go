package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"torre-response-parser/pkg/jobs"
)

const (
	FILE_PATH = "/home/carlos/dev/torre/"
)

func main() {
	file := FILE_PATH + "jobs.json"
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	bytesJobs, _ := ioutil.ReadAll(jsonFile)

	resultados, err := jobs.GetResults(bytesJobs)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(resultados); i++ {
		fmt.Println("Job N-" + strconv.Itoa(i+1) + ": " + resultados[i].Objective)
	}

}
