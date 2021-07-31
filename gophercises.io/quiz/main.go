package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Question struct{
	Ask string
	Answere string
}

type Quiz struct{
	Questions []Question
	Correct int
	Total int
}

func ParseQuestions(text [][]string) ([]Question, error){
	questions := make([]Question, len(text))

	for i ,t := range text{
		questions[i] = Question{
			Ask: t[0], 
			Answere: t[1],
		}
	}

	return questions,nil
}

func main() {
	filename := flag.String("csv", "problems.csv", "Problems csv file in format <question>,<answere>")
	flag.Parse()

	f, err := os.Open(*filename)
	validate(err, "Csv file not found!")

	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)

	lines, err := r.ReadAll()
	validate(err, "Csv file: read error")

	parsedLines, err := ParseQuestions(lines)
	validate(err, "Can't parse data from csv file")

	quiz := Quiz{Questions: parsedLines, Total: len(parsedLines)}

	for _, question := range quiz.Questions{
		fmt.Printf("Question: %s=", question.Ask)
		scanner.Scan()
		validate(scanner.Err(), "read user input error")
		text := scanner.Text()
		if text == question.Answere {
			quiz.Correct++
		}
	}

	fmt.Printf("Test Completed: Total Questions: %v, Correct Answeres: %v", quiz.Total, quiz.Correct)
}

func validate(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err.Error())
	}
}
