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
	csvFileName := getCsvFileName()
	questionAnswerMap := readCsvFileIntoMap(*csvFileName)
	correctAnswersCount, questionsCount := playTheGame(questionAnswerMap)

	fmt.Println("Your score is", correctAnswersCount, "out of", questionsCount)
}

func playTheGame(questionToAnswerMap map[string]string) (int, int) {
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
	return correctAnswersCount, questionsCount
}

func getCsvFileName() *string {
	fileNamePtr := flag.String("fileName", "problems.csv", "this is filename to read")
	flag.Parse()
	fmt.Println("Using file: ", *fileNamePtr)
	return fileNamePtr
}

func readCsvFileIntoMap(csvFileName string) map[string]string {
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
