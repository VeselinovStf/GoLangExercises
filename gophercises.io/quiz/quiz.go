package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Questions []Question
	Correct   int
}

type Question struct {
	Ask     string
	Answere string
}

// Run quiz
func (q *Quiz) Run(timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	scanner := bufio.NewScanner(os.Stdin)

	defer timer.Stop()
	for i, question := range q.Questions {
		fmt.Printf("Question %v: %s = ", i+1, question.Ask)

		answere := make(chan string)

		go func() {
			scanner.Scan()
			if scanner.Err() != nil {
				log.Fatalf("%s: %s", "Can't read user input correctly: ", scanner.Err().Error())
			}

			text := strings.TrimSpace(scanner.Text())
			answere <- text
		}()

		select {
		case <-timer.C:
			fmt.Println("Time limit! Answer Faster next time!")
			return
		case text := <-answere:
			if strings.EqualFold(text, question.Answere) {
				q.Correct++
			}
		}

	}

}

// ParseQuestions Parses quiz questions
func (q *Quiz) ParseQuestions(text [][]string) {
	questions := make([]Question, len(text))

	for i, t := range text {
		questions[i] = Question{
			Ask:     t[0],
			Answere: strings.TrimSpace(t[1]),
		}
	}

	q.Questions = questions
}

// ShufleQuestions Shufles the questions
func (q *Quiz) ShufleQuestions() {
	rand.Seed(time.Now().Unix())
	shufled := make([]Question, 0)

	for len(q.Questions) > 0 {
		r := rand.Intn(len(q.Questions))
		shufled = append(shufled, q.Questions[r])
		q.Questions = append(q.Questions[:r], q.Questions[r+1:]...)
	}

	q.Questions = shufled
}

// PrintResult print final result to client
func (q *Quiz) PrintResult() {
	fmt.Printf("Test Completed: Total Questions: %v, Correct Answeres: %v", len(q.Questions), q.Correct)
}
