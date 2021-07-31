package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := flag.String("csv", "problems.csv", "Problems csv file")
	flag.Parse()

	f, err := os.Open(*filename)
	validate(err, "problems csv file not found")

	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)

	totalQuestions := 0
	correct := 0
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		validate(err, "csv file read error")

		question := line[0]
		answere, err := strconv.Atoi(line[1])

		validate(err, "can't parse answere")

		fmt.Printf("Question: %s=", question)

		scanner.Scan()

		validate(scanner.Err(), "read input error")

		text := scanner.Text()

		userAnswere, err := strconv.Atoi(text)

		validate(err, "can't parse answere")

		if userAnswere == answere {
			correct++
		}

		totalQuestions++
	}

	fmt.Printf("Test Completed: Total Questions: %v, Correct Answeres: %v", totalQuestions, correct)
}

func validate(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err.Error())
	}
}
