package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return result
}

func main() {
	fileName := flag.String("csv", "problems.csv", "csv file with question and answer format")
	timeLimit := flag.Int("limit", 30, "Time limit in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", *fileName, err)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read csv file: %v\n", err)
		os.Exit(1)
	}

	problems := parseLines(lines)
	score := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerChan := make(chan string)
		timerChan := time.After(time.Duration(*timeLimit) * time.Second)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timerChan:
			fmt.Println("Timeout")
		case answer := <-answerChan:
			if p.answer == answer {
				fmt.Println("Correct")
				score++
			} else {
				fmt.Println("Wrong")
			}
		}
	}
	fmt.Printf("You Score is %d out of %v", score, len(lines))
}
