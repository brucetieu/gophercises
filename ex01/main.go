package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	question string
	answer   string
}

func main() {

	inputPtr := flag.String("csv", "problem.csv", "a csv file in the format of question,answer")
	timerPtr := flag.String("time", "15", "Time limit for taking the quiz")
	shufflePtr := flag.String("shuffle", "false", "Shuffle the quiz order each time it is run.")
	flag.Parse()

	csvFile, err := os.Open(*inputPtr)
	if err != nil {
		log.Fatalln("Couldn't open file: ", err)
	}

	r := csv.NewReader(csvFile)

	fmt.Println("Press Enter to start the timer")
	var key string
	fmt.Scanln(&key)

	num, err := strconv.Atoi(*timerPtr)

	records, err := r.ReadAll()

	parsedRecords := parse(records)
	if *shufflePtr == "true" {
		shuffleQuestions(parsedRecords)
	}
	answers, correctAns := collectAnswers(parsedRecords, num)
	calculateScore(answers, correctAns, len(parsedRecords))

}

func parse(records [][]string) []Record {
	var parsedRecords = make([]Record, len(records))
	for i, record := range records {
		parsedRecords[i] = Record{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return parsedRecords
}

func collectAnswers(parsedRecords []Record, duration int) (map[int]int, int) {
	correctAns := 0
	var answers = make(map[int]int)

	timeLimit := time.NewTimer(time.Duration(duration) * time.Second)

	for i, record := range parsedRecords {
		ch := make(chan string)

		go func() {
			var ans string
			fmt.Printf("Question %d: %s ", i+1, record.question)
			fmt.Scanln(&ans)
			ch <- ans
		}()

		select {
		case result := <-ch:
			if result == record.answer {
				correctAns++
				answers[i] = 1
			} else {
				answers[i] = -1
			}
		case <-timeLimit.C:
			fmt.Println("\nTime's up!")
			return answers, correctAns
		}

	}

	return answers, correctAns
}

func calculateScore(answers map[int]int, correctAns int, totalQuestions int) {
	var percentageScore float32 = float32(correctAns) / float32(totalQuestions) * 100
	fmt.Printf("Your score: %d/%d\n", correctAns, totalQuestions)
	fmt.Printf("With percent: %f\n", percentageScore)
	for key, val := range answers {
		if val != -1 {
			fmt.Printf("You got Question %d right\n", key+1)
		} else {
			fmt.Printf("You got Question %d wrong \n", key+1)
		}
	}
}

func shuffleQuestions(records []Record) {
	rand.Seed(time.Now().UnixNano())
	n := len(records)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		records[i], records[j] = records[j], records[i]
	}

}
