package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	freshIDRanges, availableIngredientIDs := loadInput("2025/day_5/input.txt")
	countFreshIngredients(freshIDRanges, availableIngredientIDs)
	countUniqueFreshIDs(freshIDRanges)
}

func countFreshIngredients(freshIDRanges [][2]int, ids []int) {
	// brute force approach
	var freshIDCount int
	for _, id := range ids {
		for _, freshIDRange := range freshIDRanges {
			if id >= freshIDRange[0] && id <= freshIDRange[1] {
				freshIDCount++
				break
			}
		}
	}

	fmt.Printf("Number of fresh ingredients: %d\n", freshIDCount)
}

func countUniqueFreshIDs(freshIDRanges [][2]int) {
	nonOverlappingFreshIDs := [][2]int{freshIDRanges[0]}
	for _, idRange := range freshIDRanges[1:] {
		var newNonOverlappingIDS [][2]int
		var newRange = idRange
		// recreate a slice of unique ID ranges. If there is any overlap, combine the ranges.
		for _, nonOverlappingRange := range nonOverlappingFreshIDs {
			// check if there is any overlap
			if idRange[0] > nonOverlappingRange[1] || idRange[1] < nonOverlappingRange[0] {
				// no overlap
				newNonOverlappingIDS = append(newNonOverlappingIDS, nonOverlappingRange)
				continue
			}

			// some overlap, set newRange to the combined range
			if nonOverlappingRange[0] < newRange[0] {
				newRange[0] = nonOverlappingRange[0]
			}
			if nonOverlappingRange[1] > newRange[1] {
				newRange[1] = nonOverlappingRange[1]
			}
		}

		newNonOverlappingIDS = append(newNonOverlappingIDS, newRange)
		nonOverlappingFreshIDs = newNonOverlappingIDS
	}

	var totalFreshIDCount int
	for _, idRange := range nonOverlappingFreshIDs {
		totalFreshIDCount += idRange[1] - idRange[0] + 1
	}

	fmt.Printf("Number of unique fresh IDs: %d\n", totalFreshIDCount)
}

func loadInput(filepath string) ([][2]int, []int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var freshIDRanges [][2]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		}
		startIndex, endIndex, ok := strings.Cut(data, "-")
		if !ok {
			panic("could not split ID range")
		}

		startIndexInt, err := strconv.Atoi(startIndex)
		if err != nil {
			panic(err)
		}
		endIdexInt, err := strconv.Atoi(endIndex)
		if err != nil {
			panic(err)
		}

		freshIDRanges = append(freshIDRanges, [2]int{startIndexInt, endIdexInt})
	}

	var availableIngredientIDs []int
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		}

		intID, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}

		availableIngredientIDs = append(availableIngredientIDs, intID)
	}

	return freshIDRanges, availableIngredientIDs
}
