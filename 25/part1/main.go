package main

import (
	"bufio"
	"fmt"
	"os"
)

const east = 1
const south = 2

func parseInput(filename string) [][]int {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	cucumbers := [][]int{}
	for s.Scan() {
		line := s.Text()
		cucumberLine := []int{}
		for _, c := range line {
			if c == '>' {
				cucumberLine = append(cucumberLine, east)
			} else if c == 'v' {
				cucumberLine = append(cucumberLine, south)
			} else {
				cucumberLine = append(cucumberLine, 0)
			}
		}
		cucumbers = append(cucumbers, cucumberLine)
	}

	return cucumbers
}

func stepSouth(cucumbers [][]int) [][]int {
	newCucumbers := make([][]int, len(cucumbers))
	for i := range cucumbers {
		newCucumbers[i] = make([]int, len(cucumbers[0]))
	}

	for col := 0; col < len(cucumbers[0]); col++ {
		for row := 0; row < len(cucumbers); row++ {
			if cucumbers[row][col] == south {
				newRow := (row + 1) % len(cucumbers)
				if cucumbers[newRow][col] == 0 {
					newCucumbers[newRow][col] = south
					row++
					continue
				}
			}
			newCucumbers[row][col] = cucumbers[row][col]
		}
	}

	return newCucumbers
}

func stepEast(cucumbers [][]int) [][]int {
	newCucumbers := make([][]int, len(cucumbers))
	for i := range cucumbers {
		newCucumbers[i] = make([]int, len(cucumbers[0]))
	}

	for row := range cucumbers {
		for col := 0; col < len(cucumbers[0]); col++ {
			if cucumbers[row][col] == east {
				newCol := (col + 1) % len(cucumbers[0])
				if cucumbers[row][newCol] == 0 {
					newCucumbers[row][newCol] = east
					col++
					continue
				}
			}
			newCucumbers[row][col] = cucumbers[row][col]
		}
	}

	return newCucumbers
}

func step(cucumbers [][]int) [][]int {
	newCucumbers := stepEast(cucumbers)
	newCucumbers = stepSouth(newCucumbers)
	return newCucumbers
}

func print(cucumbers [][]int) {
	for _, row := range cucumbers {
		for _, c := range row {
			if c == east {
				fmt.Printf(">")
			} else if c == south {
				fmt.Printf("v")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func compareCucumbers(c1, c2 [][]int) bool {
	for row := range c1 {
		for col := range c1[0] {
			if c1[row][col] != c2[row][col] {
				return false
			}
		}
	}
	return true
}

func main() {
	cucumbers := parseInput("input.txt")
	count := 0
	for {
		newCucumbers := step(cucumbers)
		count++
		fmt.Println("STEP", count)
		if compareCucumbers(newCucumbers, cucumbers) {
			fmt.Println("SAME")
			break
		}
		cucumbers = newCucumbers
	}
}
