/*
Program that reads in a quiz provided via a CSV file  and then give the quiz to a user keeping track
of how many questions they get right and how many they get incorrect.

The default csv file is hardcoded in flags ,and is customizable via a flag.
The default question time limit is 30 seconds,and is customizable via a flag.
The default questions are in csv order but can be shufled via a flag.
*/
package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

// Flags params
var (
	fileName  string
	timeLimit int
	shufle    bool
)

// Initializes Flags
func init() {
	flag.StringVar(&fileName, "csv", "problems.csv", "Problems csv file in format <question>,<answere>")
	flag.IntVar(&timeLimit, "time", 30, "Question time limit in seconds")
	flag.BoolVar(&shufle, "shufle", false, "Shuffle the quiz order.")
	flag.Parse()
}

func main() {
	lines, err := readCSV(fileName)
	if err != nil {
		log.Fatalf("%s: %s", "Can't read input csv file: ", err.Error())
	}

	quiz := Quiz{}

	quiz.ParseQuestions(lines)

	if shufle {
		quiz.ShufleQuestions()
	}

	quiz.Run(timeLimit)

	quiz.PrintResult()
}

// Reads .csv file
func readCSV(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
