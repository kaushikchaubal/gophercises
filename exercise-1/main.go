package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fileNamePtr := flag.String("fileName", "problems.csv", "this is filename to read")
	flag.Parse()
	fmt.Println("Using file: ", *fileNamePtr)

	questionToAnswerMap := readCSVtoMap(*fileNamePtr)
	correctAnswersCount := 0
	questionsCount := 0

	reader := bufio.NewReader(os.Stdin)
	for question := range questionToAnswerMap {
		questionsCount++
		fmt.Println("Your next question is ", question)

		userAnswer, _ := reader.ReadString('\n')

		correctAnswer := questionToAnswerMap[question]

		if strings.TrimSpace(userAnswer) == correctAnswer {
			fmt.Println("You are a rockstar!")
			correctAnswersCount++

		} else {
			fmt.Println("You suck!")
		}

	}
	fmt.Println("Your score is ", correctAnswersCount, " out of ", questionsCount)
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
		questionToAnswerMap[record[0]] = strings.TrimSpace(record[1])
	}

	return questionToAnswerMap
}
