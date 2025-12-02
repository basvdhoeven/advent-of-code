package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2025/day_2/input.txt")
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	dataString = strings.TrimSuffix(dataString, "\n")

	ranges := strings.Split(dataString, ",")

	var repeatedIDsSum int

	for _, singleRange := range ranges {
		first, last, found := strings.Cut(singleRange, "-")
		if !found {
			panic("invalid input")
		}

		firstNum, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}
		lastNum, err := strconv.Atoi(last)
		if err != nil {
			panic(err)
		}

		for i := firstNum; i <= lastNum; i++ {
			IDstring := strconv.Itoa(i)
			length := len(IDstring)
			if len(IDstring)%2 != 0 {
				// Odd amount of numbers
				continue
			}

			if IDstring[:length/2] == IDstring[length/2:] {
				repeatedIDsSum += i
			}
		}
	}

	fmt.Println(fmt.Sprintf("invalid IDs sum: %d", repeatedIDsSum))
}
