package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type question struct {
	content  string
	solution string
}

func main() {
	score := 0
	var quiz []question

	file, _ := os.Open("problems.csv")
	result := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := result.Read()
		if err == io.EOF {
			break
		}
		quiz = append(quiz, question{
			content:  record[0],
			solution: record[1],
		})

	}

	for _, val := range quiz {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Hi! What is the value of ", val.content, "?")

		for scanner.Scan() {
			answer := scanner.Text()
			validateAnswer(answer, val.solution, &score)
			break
		}

	}
	fmt.Println("the final score is: ", score)
}

func validateAnswer(answer string, solution string, score *int) {
	if answer == solution {
		fmt.Println("Correct!")
		*score = *score + 1

	} else {
		fmt.Println("Incorrect!")
	}

}
