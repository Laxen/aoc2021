package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

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

func findPatternsFromSegmentCount(patterns []string, count int) []string {
	ret := []string{}
	for _, pattern := range patterns {
		if len(pattern) == count {
			ret = append(ret, pattern)
		}
	}
	return ret
}

func subtractPattern(p1 string, p2 string) string {
	ret := ""
	for _, l1 := range p1 {
		add := true
		for _, l2 := range p2 {
			if l1 == l2 {
				add = false
				break
			}
		}

		if add {
			ret += string(l1)
		}
	}

	return ret
}

func getLetterMap(pattern []string) map[string]rune {
	lettermap := map[rune]string{}
	numbermap := map[int]string{}

	numbermap[1] = findPatternsFromSegmentCount(pattern, 2)[0]
	numbermap[4] = findPatternsFromSegmentCount(pattern, 4)[0]
	numbermap[7] = findPatternsFromSegmentCount(pattern, 3)[0]
	numbermap[8] = findPatternsFromSegmentCount(pattern, 7)[0]

	lettermap['a'] = subtractPattern(numbermap[7], numbermap[1])

	// g = 9 - 4 - a
	pat := findPatternsFromSegmentCount(pattern, 6)
	for _, n := range pat {
		tmp := subtractPattern(n, numbermap[4])
		tmp = subtractPattern(tmp, lettermap['a'])
		if len(tmp) == 1 {
			lettermap['g'] = tmp
			numbermap[9] = n
			break
		}
	}

	// e = 2 - 4 - a - g
	pat = findPatternsFromSegmentCount(pattern, 5)
	for _, n := range pat {
		tmp := subtractPattern(n, numbermap[4])
		tmp = subtractPattern(tmp, lettermap['a'])
		tmp = subtractPattern(tmp, lettermap['g'])
		if len(tmp) == 1 {
			lettermap['e'] = tmp
			numbermap[2] = n
			break
		}
	}

	// b = 0 - 1 - a - g - e
	pat = findPatternsFromSegmentCount(pattern, 6)
	for _, n := range pat {
		tmp := subtractPattern(n, numbermap[1])
		tmp = subtractPattern(tmp, lettermap['a'])
		tmp = subtractPattern(tmp, lettermap['g'])
		tmp = subtractPattern(tmp, lettermap['e'])
		if len(tmp) == 1 {
			lettermap['b'] = tmp
			numbermap[0] = n
			break
		}
	}

	// f = 1 - 2
	lettermap['f'] = subtractPattern(numbermap[1], numbermap[2])

	// d = 8 - 0
	lettermap['d'] = subtractPattern(numbermap[8], numbermap[0])

	// c = 1 - f
	lettermap['c'] = subtractPattern(numbermap[1], lettermap['f'])

	reverseLetterMap := map[string]rune{}
	for key, val := range lettermap {
		reverseLetterMap[val] = key
	}

	return reverseLetterMap
}

func sortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func decode(lettermap map[string]rune, pattern string) int {
	decoded := ""
	for _, letter := range pattern {
		decoded += string(lettermap[string(letter)])
	}

	decoded = sortString(decoded)

	switch decoded {
	case "abcefg":
		return 0
	case "cf":
		return 1
	case "acdeg":
		return 2
	case "acdfg":
		return 3
	case "bcdf":
		return 4
	case "abdfg":
		return 5
	case "abdefg":
		return 6
	case "acf":
		return 7
	case "abcdefg":
		return 8
	case "abcdfg":
		return 9
	}

	return 0
}

func pow(a int, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

func main() {
	patterns, outputs := parseInput("input.txt")

	sum := 0
	for line, _ := range patterns {
		lettermap := getLetterMap(patterns[line])
		subsum := 0
		for i := 0; i < 4; i++ {
			d := decode(lettermap, outputs[line][i])
			subsum += d * pow(10, 3-i)
		}
		sum += subsum
	}
	fmt.Println(sum)
}
