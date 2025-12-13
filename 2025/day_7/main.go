package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := loadInput("2025/day_7/input.txt")

	calculateNumberOfSplits(input)
	calculateNumberOfTimelines(input)
}

func calculateNumberOfTimelines(input [][]rune) {
	timeLinesArray := make([][]int, len(input))
	for i := range timeLinesArray {
		timeLinesArray[i] = make([]int, len(input[0]))
	}

	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input[0]); j++ {
			switch input[i][j] {
			case 'S':
				// starting point
				timeLinesArray[i+1][j] = 1
			case '^':
				// split
				timeLinesArray[i+1][j-1] += timeLinesArray[i-1][j]
				timeLinesArray[i+1][j+1] += timeLinesArray[i-1][j]
			default:
				timeLinesArray[i+1][j] += timeLinesArray[i][j]
			}
		}
	}

	var totalTimelines int
	for _, val := range timeLinesArray[len(input)-1] {
		totalTimelines += val
	}
	fmt.Printf("Total timelines: %d\n", totalTimelines)
}

func calculateNumberOfSplits(input [][]rune) {
	var splits int
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == 'S' {
				// starting point
				input[i+1][j] = '|'
			}

			if input[i][j] == '|' {
				if input[i+1][j] == '^' {
					splits++
					input[i+1][j-1] = '|'
					input[i+1][j+1] = '|'
				} else {
					input[i+1][j] = '|'
				}
			}
		}
	}
	fmt.Printf("Number of splits: %d\n", splits)
}

func loadInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allData [][]rune
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		}
		allData = append(allData, []rune(data))
	}

	return allData
}
