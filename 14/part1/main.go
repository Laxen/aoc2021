package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(filename string) (map[string]string, string) {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	s.Scan()
	poly := s.Text()
	s.Scan()

	m := map[string]string{}
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, " -> ")
		m[split[0]] = split[1]
	}
	return m, poly
}

func work(polyMap map[string]string, poly string) string {
	newPoly := ""
	for i := 0; i < len(poly)-1; i++ {
		pair := poly[i : i+2]
		newPoly += string(poly[i])
		if polyMap[pair] != "" {
			newPoly += polyMap[pair]
		}
	}
	newPoly += string(poly[len(poly)-1])
	return newPoly
}

func findMostLeastCommon(poly string) (int, int) {
	count := map[rune]int{}
	for _, c := range poly {
		count[c]++
	}

	most := 0
	least := 10000
	for _, v := range count {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most, least
}

func main() {
	polyMap, poly := parseInput("input.txt")
	for i := 0; i < 10; i++ {
		fmt.Println("Step:", i)
		poly = work(polyMap, poly)
	}
	most, least := findMostLeastCommon(poly)
	fmt.Println(most - least)
}
