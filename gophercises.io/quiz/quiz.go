package main

import (
	"bufio"
	"fmt"
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

func (q *Quiz) Run(scanner *bufio.Scanner, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, question := range q.Questions {
		fmt.Printf("Question %v: %s = ", i+1, question.Ask)

		answere := make(chan string)

		go func() {
			scanner.Scan()
			validate(scanner.Err(), "read user input error")
			text := strings.TrimSpace(scanner.Text())
			answere <- text
		}()

		select {
		case <-timer.C:
			fmt.Println("Time limit! Answer Faster next time!")
			return
		case text := <-answere:
			if text == question.Answere {
				q.Correct++
			}
		}

	}
}

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

func (q *Quiz) PrintResult() {
	fmt.Printf("Test Completed: Total Questions: %v, Correct Answeres: %v", len(q.Questions), q.Correct)
}
