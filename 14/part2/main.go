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
	var newPoly strings.Builder
	for i := 0; i < len(poly)-1; i++ {
		pair := poly[i : i+2]
		newPoly.WriteByte(poly[i])
		if polyMap[pair] != "" {
			newPoly.WriteString(polyMap[pair])
		}
	}
	newPoly.WriteByte(poly[len(poly)-1])
	return newPoly.String()
}

func countLetters(poly string) map[string]int {
	count := map[string]int{}
	for _, c := range poly {
		count[string(c)]++
	}
	return count
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

func polyMapSeq(polyMap map[string]string, steps int) (map[string]string, map[string]map[string]int) {
	ret := map[string]string{}
	count := map[string]map[string]int{}

	for pair := range polyMap {
		res := pair
		fmt.Println("Pair:", pair)
		for i := 0; i < steps; i++ {
			res = work(polyMap, res)
		}
		ret[pair] = res[:len(res)-1]
		count[pair] = countLetters(ret[pair])
	}

	return ret, count
}

func polyMapSeqCount(count map[string]map[string]int, poly string, letter string) int {
	letterCount := 0

	for i := 0; i < len(poly)-1; i++ {
		pair := poly[i : i+2]
		letterCount += count[pair][letter]
	}
	if string(poly[len(poly)-1]) == letter {
		letterCount++
	}

	return letterCount
}

func polyMapSeqCount2(count map[string]map[string]int, poly string, letters []string) map[string]int {
	letterCount := map[string]int{}

	for i := 0; i < len(poly)-1; i++ {
		pair := poly[i : i+2]
		for _, letter := range letters {
			letterCount[letter] += count[pair][letter]
		}
	}

	letterCount[string(poly[len(poly)-1])]++

	return letterCount
}

func findUniqueLetters(polyMap map[string]string) []string {
	list := []string{}
	for _, v := range polyMap {
		cont := false
		for _, e := range list {
			if v == e {
				cont = true
				break
			}
		}

		if cont {
			continue
		} else {
			list = append(list, v)
		}
	}

	return list
}

func findMostLeastCommon2(count map[string]int) (int, int) {
	most := 0
	least := 184467440737095516

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
	steps := 20
	_, cnt := polyMapSeq(polyMap, steps)

	for i := 0; i < steps; i++ {
		fmt.Println("Step:", i)
		poly = work(polyMap, poly)
	}

	letters := findUniqueLetters(polyMap)
	fmt.Println(letters)

	nCnt2 := polyMapSeqCount2(cnt, poly, letters)
	fmt.Println(nCnt2)

	most, least := findMostLeastCommon2(nCnt2)
	fmt.Println(most, least)
	fmt.Println(most - least)
}
