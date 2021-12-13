package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([][]int, int, int) {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	xMax := 0
	yMax := 0

	dots := [5000][5000]int{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		dots[y][x] = 1

		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	xFold := 0
	yFold := 0
	for s.Scan() {
		line := s.Text()
		t := strings.Split(line, " ")
		fold := strings.Split(t[2], "=")
		if fold[0] == "x" {
			x, _ := strconv.Atoi(fold[1])
			xFold = x
		} else if fold[0] == "y" {
			y, _ := strconv.Atoi(fold[1])
			yFold = y
		}
	}

	ret := [][]int{}
	for y := 0; y <= yMax; y++ {
		list := []int{}
		for x := 0; x <= xMax; x++ {
			list = append(list, dots[y][x])
		}
		ret = append(ret, list)
	}

	return ret, xFold, yFold
}

func countDots(dots [][]int) (r int) {
	for y := range dots {
		for x := range dots[0] {
			if dots[y][x] > 0 {
				r++
			}
		}
	}

	return r
}

func yFold(dots [][]int, yFold int) [][]int {
	rows := len(dots)

	for y := 0; y < yFold; y++ {
		for x := range dots[0] {
			dots[y][x] += dots[rows-y-1][x]
		}
	}

	return dots[:yFold][:]
}

func xFold(dots [][]int, xFold int) [][]int {
	cols := len(dots[0])

	for y := range dots {
		for x := 0; x < xFold; x++ {
			dots[y][x] += dots[y][cols-x-1]
		}
	}

	fold := [][]int{}
	for _, y := range dots {
		fold = append(fold, y[:xFold])
	}
	return fold
}

func printDots(dots [][]int) {
	for y := range dots {
		for x := range dots[0] {
			if dots[y][x] > 0 {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	dots, _, _ := parseInput("input.txt")
	dots = xFold(dots, 655)
	dots = yFold(dots, 447)
	dots = xFold(dots, 327)
	dots = yFold(dots, 223)
	dots = xFold(dots, 163)
	dots = yFold(dots, 111)
	dots = xFold(dots, 81)
	dots = yFold(dots, 55)
	dots = xFold(dots, 40)
	dots = yFold(dots, 27)
	dots = yFold(dots, 13)
	dots = yFold(dots, 6)

	printDots(dots)
}
