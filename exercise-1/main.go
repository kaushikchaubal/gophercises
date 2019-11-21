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
	"time"
)

func main() {
	csvFileName, timerDuration := getDataFromFlags()
	questionAnswerMap, questionsCount := readCsvFileIntoMap(csvFileName)
	correctAnswersCount := playTheGame(questionAnswerMap, timerDuration)

	fmt.Println("Your score is", correctAnswersCount, "out of", questionsCount)
}

func playTheGame(questionToAnswerMap map[string]string, timerDuration time.Duration) int {
	correctAnswersCount := 0

	reader := bufio.NewReader(os.Stdin)

	channel := make(chan string)
	timer := time.NewTimer(timerDuration)

	go func() {
		<-timer.C
		fmt.Println("Time up!")
		channel <- "showResult"
	}()

	go func() {
		for question := range questionToAnswerMap {
			interactiveQA(&correctAnswersCount, question, questionToAnswerMap, reader)
		}

		timer.Stop()
		channel <- "showResult"
	}()

	<-channel
	return correctAnswersCount
}

func interactiveQA(correctAnswersCount *int, question string, questionToAnswerMap map[string]string, reader *bufio.Reader) {
	fmt.Println("Your next question is ", question)

	userAnswer, _ := reader.ReadString('\n')

	correctAnswer := questionToAnswerMap[question]

	if strings.TrimSpace(userAnswer) == correctAnswer {
		fmt.Println("You are a rockstar!")
		*correctAnswersCount++
	} else {
		fmt.Println("You suck!")
	}
}

func getDataFromFlags() (string, time.Duration) {
	fileNamePtr := flag.String("fileName", "problems.csv", "this is filename to read")
	timerDurationPtr := flag.Duration("timerDuration", 4*time.Second, "Time duration for the quiz before it times out")
	flag.Parse()

	return *fileNamePtr, *timerDurationPtr
}

func readCsvFileIntoMap(csvFileName string) (map[string]string, int) {
	questionsCount := 0
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
		questionsCount++
	}

	return questionToAnswerMap, questionsCount
}
