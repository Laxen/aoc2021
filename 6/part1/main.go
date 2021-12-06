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

type fish struct {
	timer int
}

func parseInput(filename string) []fish {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	ages := strings.Split(line, ",")

	fishes := []fish{}
	for _, age := range ages {
		age, _ := strconv.Atoi(age)
		fish := fish{age}
		fishes = append(fishes, fish)
	}

	return fishes
}

func nextDay(fishes *[]fish) {
	newFishes := []fish{}

	for i := range *fishes {
		t := (*fishes)[i].timer
		if t == 0 {
			newFishes = append(newFishes, fish{8})
			t = 6
		} else {
			t--
		}
		(*fishes)[i].timer = t
	}

	*fishes = append(*fishes, newFishes...)
}

func main() {
	fishes := parseInput("input.txt")
	for i := 0; i < 80; i++ {
		nextDay(&fishes)
	}
	fmt.Println(fishes)
	fmt.Println(len(fishes))
}
