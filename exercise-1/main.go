package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	questionToAnswerMap := readCSVtoMap("problems.csv")
	score := 0
	questions := 0

	reader := bufio.NewReader(os.Stdin)
	for question := range questionToAnswerMap {
		questions++
		fmt.Println("Your next question is ", question)

		userAnswer, _ := reader.ReadString('\n')

		correctAnswer := questionToAnswerMap[question]

		if strings.TrimSpace(userAnswer) == correctAnswer {
			fmt.Println("You are a rockstar!")
			score++

		} else {
			fmt.Println("You suck!")
		}

	}
	fmt.Println("Your score is ", score, " out of ", questions)
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
