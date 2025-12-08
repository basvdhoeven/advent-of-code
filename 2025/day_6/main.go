package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("2025/day_6/input.txt")
	// sumOfProblemSolutions(input)
	rightToLeftColumnsSolutions(input)
}

func rightToLeftColumnsSolutions(lines []string) {
	var problems []problem
	var operator string
	var numbers []int
	for i := 0; i < len(lines[0]); i++ {
		var number []rune
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != ' ' {
				if j == len(lines)-1 {
					operator = string(lines[j][i])
				} else {
					number = append(number, rune(lines[j][i]))
				}
			}
		}
		if number == nil {
			// end of problem
			problems = append(problems, problem{numbers, operator})
			numbers = nil
		} else {
			intNumber, err := strconv.Atoi(string(number))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, intNumber)
		}
	}

	problems = append(problems, problem{numbers, operator})

	var sumAllProblems int
	for _, problem := range problems {
		if problem.operator == "*" {
			solution := 1
			for _, n := range problem.numbers {
				solution *= n
			}
			sumAllProblems += solution
		}
		if problem.operator == "+" {
			for _, n := range problem.numbers {
				sumAllProblems += n
			}
		}
	}

	fmt.Printf("Correct solution for all problems: %d\n", sumAllProblems)
}

func sumOfProblemSolutions(data []string) {

	columns := len(strings.Fields(data[0]))

	problems := make([]problem, columns)

	for i, row := range data {
		entries := strings.Fields(row)
		for j, entry := range entries {
			if i < 4 {
				numberInt, err := strconv.Atoi(entry)
				if err != nil {
					panic(err)
				}
				problems[j].numbers = append(problems[j].numbers, numberInt)
			} else {
				problems[j].operator = entry
			}
		}
	}

	var sumAllProblems int
	for _, problem := range problems {
		switch problem.operator {
		case "*":
			sumAllProblems += problem.numbers[0] * problem.numbers[1] * problem.numbers[2] * problem.numbers[3]
		case "+":
			sumAllProblems += problem.numbers[0] + problem.numbers[1] + problem.numbers[2] + problem.numbers[3]
		}
	}

	fmt.Printf("Sum of all problem solutions: %d\n", sumAllProblems)
}

type problem struct {
	numbers  []int
	operator string
}

func loadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allData []string
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		}
		allData = append(allData, data)
	}

	return allData
}
