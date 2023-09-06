package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

type problem struct {
	q string
	a string
}

// Function to colorize the text outputs
func colorize(color Color, message string) {
	fmt.Printf("%s%s%s", string(color), message, string(ColorReset))
}

// Parsing lines
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	banner := "Welcome to nihon-Go quiz game! \nようこそnihon-Goクイズゲームへ!\n"
	colorize(ColorBlue, banner)
	csvFilename := flag.String("csv", "problems.csv", "Provide a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		msg := fmt.Sprintf("Failed to open file: <%s>, Please try again!", *csvFilename)
		colorize(ColorRed, msg)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		colorize(ColorRed, "Failed to read CSV file!")
		os.Exit(1)
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Question #%d : %s\n=> ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		answer = strings.ToUpper(answer)
		if answer == p.a {
			colorize(ColorGreen, "Correct Answer!\n")
			correct++
		} else {
			colorize(ColorRed, "Wrong Answer!\n")
		}
	}

	youScored := fmt.Sprintf("You scored %d out of %d. \n", correct, len(problems))
	colorize(ColorBlue, youScored)
}
