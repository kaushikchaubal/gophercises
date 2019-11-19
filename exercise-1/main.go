package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	readCSVtoMap("problems.csv")
	fmt.Println(questionToAnswerMap)
}

func readCSVtoMap(csvFileName string) map[string]string {
	csvfile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	reader := csv.NewReader(csvfile)
	questionToAnswerMap := make(map[string]string)

	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questionToAnswerMap[record[0]] = record[1]
	}

	return questionToAnswerMap
}
