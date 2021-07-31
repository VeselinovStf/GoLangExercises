package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
)


func main() {
	filename := flag.String("csv", "problems.csv", "Problems csv file in format <question>,<answere>")
	timeLimit := flag.Int("time", 30, "Question time limit in seconds")
	shufle := flag.Bool("shufle", false, "Shuffle the quiz order.")

	flag.Parse()

	f, err := os.Open(*filename)
	validate(err, "Csv file not found!")

	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)

	lines, err := r.ReadAll()
	validate(err, "Csv file: read error")

	quiz := Quiz{}

	quiz.ParseQuestions(lines)

	if *shufle {
		quiz.ShufleQuestions()
	}

	quiz.Run(scanner, *timeLimit)

	quiz.PrintResult()
}

func validate(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err.Error())
	}
}
