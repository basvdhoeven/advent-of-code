package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2025/day_1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	dial := 50 // starting number
	var zeroPointings, zeroPassings int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := string(scanner.Text())

		rotation, err := strconv.Atoi(line[1:len(line)])
		if err != nil {
			panic(err)
		}

		// Moving the dial left.
		if strings.HasPrefix(line, "L") {
			zeroPassings += dial/100 - (dial-rotation)/100
			dial -= rotation
		}

		// Moving the dial right.
		if strings.HasPrefix(line, "R") {
			zeroPassings += (dial+rotation)/100 - dial/100
			dial += rotation
		}

		// check if dial is at zero.
		if dial % 100 == 0 {
			zeroPointings++
		}
	}

	fmt.Println(fmt.Sprintf("Zero pointings after rotating: %d", zeroPointings))
	fmt.Println(fmt.Sprintf("Zero passings: %d", zeroPassings))
}


