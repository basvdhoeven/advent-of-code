package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2025/day_3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalJoltage int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := scanner.Text()

		var largestJoltage int
		for i := range line {
			for j := i+1; j < len(line); j++ {
				joltage, err := strconv.Atoi(string(line[i]) + string(line[j]))
				if err != nil {
					panic(err)
				}
				if joltage > largestJoltage {
					largestJoltage = joltage
				}
			}
		}

		totalJoltage += largestJoltage
		fmt.Printf("line: %s\n", line)
		fmt.Printf("largest joltage: %d\n", largestJoltage)
	}
	fmt.Printf("Total joltage: %d\n", totalJoltage)
}
