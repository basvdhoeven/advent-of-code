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

			for j := 1; j <= length/2; j++ {
				if length%j != 0 {
					continue
				}
				count := strings.Count(IDstring, IDstring[:j])
				if count == length/j {
					repeatedIDsSum += i
					break
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("invalid IDs sum: %d", repeatedIDsSum))
}
