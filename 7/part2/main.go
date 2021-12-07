package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, int) {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	pos_str := strings.Split(line, ",")

	max := 0
	ret := []int{}
	for _, pos := range pos_str {
		pos, _ := strconv.Atoi(pos)
		ret = append(ret, pos)
		if pos > max {
			max = pos
		}
	}

	return ret, max
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func smallest(m map[int]int) int {
	ret := m[0]
	for _, e := range m {
		if e < ret {
			ret = e
		}
	}
	return ret
}

func fuelCost(a int, b int) int {
	n := abs(a - b)
	return n * (n + 1) / 2
}

func calcFuelCost(pos []int, max_pos int) map[int]int {
	costs := map[int]int{}
	for _, p := range pos {
		for i := 0; i < max_pos; i++ {
			costs[i] += fuelCost(p, i)
		}
	}

	return costs
}

func main() {
	pos, max := parseInput("input.txt")
	costs := calcFuelCost(pos, max)
	cheapest_cost := smallest(costs)
	fmt.Println(cheapest_cost)
}
