package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringListToInt(list []string) []int {
	ret := []int{}
	for _, e := range list {
		ne, _ := strconv.Atoi(e)
		ret = append(ret, ne)
	}
	return ret
}

func parseInput(filename string) [][]int {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	heightmap := [][]int{}
	for s.Scan() {
		line := s.Text()
		heights := strings.Split(line, "")
		heightmap = append(heightmap, stringListToInt(heights))
	}

	return heightmap
}

func findRiskSum(heightmap [][]int) int {
	risk := 0
	for y := range heightmap {
		for x := range heightmap[0] {
			val := heightmap[y][x]

			if x-1 >= 0 {
				if heightmap[y][x-1] <= val {
					continue
				}
			}

			if x+1 < len(heightmap[0]) {
				if heightmap[y][x+1] <= val {
					continue
				}
			}

			if y-1 >= 0 {
				if heightmap[y-1][x] <= val {
					continue
				}
			}

			if y+1 < len(heightmap) {
				if heightmap[y+1][x] <= val {
					continue
				}
			}

			risk += 1 + val
		}
	}

	return risk
}

func main() {
	heightmap := parseInput("input.txt")
	risk := findRiskSum(heightmap)
	fmt.Println(risk)
}
