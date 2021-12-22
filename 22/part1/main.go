package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setCuboid(cubes [101][101][101]int, xRange, yRange, zRange []int, value int) [101][101][101]int {
	for z := zRange[0]; z <= zRange[1]; z++ {
		for y := yRange[0]; y <= yRange[1]; y++ {
			for x := xRange[0]; x <= xRange[1]; x++ {
				cubes[z][y][x] = value
			}
		}
	}
	return cubes
}

func getNumberOfOnCubes(cubes [101][101][101]int) (ret int) {
	for z := range cubes {
		for y := range cubes[0] {
			for x := range cubes[0][0] {
				ret += cubes[z][y][x]
			}
		}
	}
	return ret
}

func parseInput(filename string) [][]int {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	rangeList := [][]int{}
	for s.Scan() {
		line := s.Text()
		onoff := strings.Split(line, " ")[0]
		ranges := strings.Split(strings.Split(line, " ")[1], ",")

		newRange := []int{}
		if onoff == "on" {
			newRange = append(newRange, 1)
		} else {
			newRange = append(newRange, 0)
		}
		for _, r := range ranges {
			s := strings.Split(r, "=")
			interval := strings.Split(s[1], "..")
			int1, _ := strconv.Atoi(interval[0])
			int2, _ := strconv.Atoi(interval[1])
			int1 += 50
			int2 += 50
			newRange = append(newRange, int1)
			newRange = append(newRange, int2)
		}
		rangeList = append(rangeList, newRange)
	}

	return rangeList
}

func main() {
	input := parseInput("input.txt")
	fmt.Println(input)

	cubes := [101][101][101]int{}

	for _, line := range input {
		if line[1] < 0 || line[2] < 0 || line[3] < 0 || line[4] < 0 || line[5] < 0 || line[6] < 0 {
			continue
		}
		if line[1] > 100 || line[2] > 100 || line[3] > 100 || line[4] > 100 || line[5] > 100 || line[6] > 100 {
			continue
		}
		cubes = setCuboid(cubes, []int{line[1], line[2]}, []int{line[3], line[4]}, []int{line[5], line[6]}, line[0])
	}

	fmt.Println(getNumberOfOnCubes(cubes))
}
