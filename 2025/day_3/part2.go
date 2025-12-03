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

		numberOfBatteries := 12
		var indexNumber int 

		var maxJoltage string

		for endOffset := numberOfBatteries;  endOffset >= 0; endOffset-- {
			var maxNumber byte
			indexNumber, maxNumber = getMaxNumber(line, indexNumber+1,  len(line)-endOffset-1)
			maxJoltage += string(maxNumber)
		}

		maxJoltageInt, _ := strconv.Atoi(maxJoltage)
		totalJoltage += maxJoltageInt
		fmt.Printf("line: %s\n", line)
		fmt.Printf("largest joltage: %d\n", maxJoltageInt)
	}
	fmt.Printf("Total joltage: %d\n", totalJoltage)
}

func getMaxNumber(s string, startIndex, endIndex int) (int, byte) {
	maxIndex := startIndex
	maxByte := s[startIndex]

	for i := startIndex; i <= endIndex; i++ {
		if s[i] > maxByte {
			maxByte = s[i]
			maxIndex = i
		}
	}

	return maxIndex, maxByte
}
