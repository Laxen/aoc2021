package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type digit struct {
	wires string
}

func parseInput(filename string) ([][]string, [][]string) {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	patterns := [][]string{}
	outputs := [][]string{}
	i := 0
	for s.Scan() {
		line := s.Text()
		p := strings.Split(line, " | ")[0]
		o := strings.Split(line, " | ")[1]

		pattern := strings.Split(p, " ")
		output := strings.Split(o, " ")

		patterns = append(patterns, pattern)
		outputs = append(outputs, output)

		i++
	}

	return patterns, outputs
}

func findNumberOfUniques(patterns [][]string) int {
	sum := 0
	for _, entry := range patterns {
		for _, pattern := range entry {
			switch len(pattern) {
			case 2:
				sum++
				break
			case 3:
				sum++
				break
			case 4:
				sum++
				break
			case 7:
				sum++
				break
			}
		}
	}

	return sum
}

func main() {
	patterns, outputs := parseInput("input.txt")
	_ = patterns
	nUniques := findNumberOfUniques(outputs)
	fmt.Println(nUniques)
}
