package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type question struct {
	question string
	answer   string
}

type quiz struct {
	questions []question
	result    int
	timeout   int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func (q *quiz) startQuiz() {
	fmt.Println("Press any button to start the quiz. ")
	scanner.Scan()

	timer := time.NewTimer(time.Duration(q.timeout) * time.Second)

	for ind, ques := range q.questions {
		fmt.Printf("Problem #%d: %s = ", ind+1, ques.question)
		ansCh := make(chan string)

		go func() {
			scanner.Scan()
			ansCh <- scanner.Text()
		}()

		select {
		case ans := <-ansCh:
			if ans == ques.answer {
				q.result++
			}
		case <-timer.C:
			return
		}
	}
}

func (q *quiz) readFileAndPrepareQuiz(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, record := range records {
		q.questions = append(q.questions, question{record[0], record[1]})
	}
}

func main() {
	q := &quiz{result: 0}
	filename := flag.String("f", "problems.csv", "")
	timeout := flag.Int("t", 1, "")

	flag.Parse()

	q.timeout = *timeout

	q.readFileAndPrepareQuiz(*filename)
	q.startQuiz()
	fmt.Printf("\nYou scored %d/%d", q.result, len(q.questions))
}
