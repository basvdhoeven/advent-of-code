package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := loadInput()
	part1(input)
	part2(input)
}

func part1(paperRolls []string) {

	lastColumn := len(paperRolls[0]) - 1
	lastRow := len(paperRolls) - 1

	var accessibleRolls int
	// loop through each location
	for i := 0; i <= lastColumn; i++ {
		for j := 0; j <= lastRow; j++ {
			if paperRolls[i][j] != 64 {
				continue // not a paper roll.
			}

			var nearbyRolls int
			if i > 0 {
				// row above exists
				if paperRolls[i-1][j] == 64 {
					nearbyRolls++ // directly above
				}
				if j < lastColumn {
					// top right location exists
					if paperRolls[i-1][j+1] == 64 {
						nearbyRolls++
					}
				}
				if j > 0 {
					// top left location exists
					if paperRolls[i-1][j-1] == 64 {
						nearbyRolls++
					}
				}
			}

			// check position to the right
			if j < lastColumn && paperRolls[i][j+1] == 64 {
				nearbyRolls++
			}
			// check position to the left
			if j > 0 && paperRolls[i][j-1] == 64 {
				nearbyRolls++
			}

			if i < lastRow {
				// row below exists
				if paperRolls[i+1][j] == 64 {
					nearbyRolls++ // directly below
				}
				if j < lastColumn {
					// bottom right location exists
					if paperRolls[i+1][j+1] == 64 {
						nearbyRolls++
					}
				}
				if j > 0 {
					// bottom left location exists
					if paperRolls[i+1][j-1] == 64 {
						nearbyRolls++
					}
				}
			}

			if nearbyRolls < 4 {
				accessibleRolls++
			}
		}
	}

	fmt.Printf("Accesible rolls: %d\n", accessibleRolls)
}

func part2(paperRolls []string) {
	lastColumn := len(paperRolls[0]) - 1
	lastRow := len(paperRolls) - 1

	var removedRolls int
	for {
		var changes bool
		// loop through each location
		for i := 0; i <= lastColumn; i++ {
			for j := 0; j <= lastRow; j++ {
				if paperRolls[i][j] != 64 {
					continue // not a paper roll.
				}

				var nearbyRolls int
				if i > 0 {
					// row above exists
					if paperRolls[i-1][j] == 64 {
						nearbyRolls++ // directly above
					}
					if j < lastColumn {
						// top right location exists
						if paperRolls[i-1][j+1] == 64 {
							nearbyRolls++
						}
					}
					if j > 0 {
						// top left location exists
						if paperRolls[i-1][j-1] == 64 {
							nearbyRolls++
						}
					}
				}

				// check position to the right
				if j < lastColumn && paperRolls[i][j+1] == 64 {
					nearbyRolls++
				}
				// check position to the left
				if j > 0 && paperRolls[i][j-1] == 64 {
					nearbyRolls++
				}

				if i < lastRow {
					// row below exists
					if paperRolls[i+1][j] == 64 {
						nearbyRolls++ // directly below
					}
					if j < lastColumn {
						// bottom right location exists
						if paperRolls[i+1][j+1] == 64 {
							nearbyRolls++
						}
					}
					if j > 0 {
						// bottom left location exists
						if paperRolls[i+1][j-1] == 64 {
							nearbyRolls++
						}
					}
				}

				if nearbyRolls < 4 {
					// remove the roll
					rowRolls := []rune(paperRolls[i])
					rowRolls[j] = '.'
					paperRolls[i] = string(rowRolls)
					changes = true
					removedRolls++
				}
			}
		}

		if !changes {
			break
		}
	}

	fmt.Printf("removed rolls: %d\n", removedRolls)
}

func loadInput() []string {
	file, err := os.Open("2025/day_4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var paperRolls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		}
		paperRolls = append(paperRolls, data)
	}

	return paperRolls
}
