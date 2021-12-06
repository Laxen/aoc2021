// Each fish creates a new fish every 7 days (after 2 day "growth" period)
// Model each fish as a number that represents the number of days until it creates a new fish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) map[int]int {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	ages := strings.Split(line, ",")

	fishes := make(map[int]int)
	for _, age := range ages {
		age, _ := strconv.Atoi(age)
		fishes[age]++
	}

	return fishes
}

func nextDay(fishes *map[int]int) {
	temp := (*fishes)[0]
	for i := 0; i < 8; i++ {
		(*fishes)[i] = (*fishes)[i+1]
	}
	(*fishes)[6] += temp
	(*fishes)[8] = temp
}

func sumFishes(fishes map[int]int) int {
	sum := 0
	for _, num := range fishes {
		sum += num
	}
	return sum
}

func main() {
	fishes := parseInput("input.txt")
	for i := 0; i < 256; i++ {
		nextDay(&fishes)
	}
	fmt.Println(fishes)
	fmt.Println(sumFishes(fishes))
}
