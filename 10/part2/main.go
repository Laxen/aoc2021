package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type stack []rune

func (s *stack) push(c rune) {
	*s = append(*s, c)
}

func (s *stack) pop() rune {
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret
}

func (s stack) toString() string {
	ret := ""
	for i := range s {
		ret += string(s[len(s)-i-1])
	}
	return ret
}

func parseInput(filename string) []string {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	ret := []string{}
	for s.Scan() {
		line := s.Text()
		ret = append(ret, line)
	}

	return ret
}

func isCloser(char rune) bool {
	switch char {
	case ')':
		fallthrough
	case ']':
		fallthrough
	case '}':
		fallthrough
	case '>':
		return true
	}
	return false
}

func getMatchingCloser(opener rune) (rune, error) {
	switch opener {
	case '(':
		return ')', nil
	case '[':
		return ']', nil
	case '{':
		return '}', nil
	case '<':
		return '>', nil
	default:
		return 'o', errors.New("No matching closer found")
	}
}

func completeLine(line string) (string, error) {
	s := stack{}
	for _, char := range line {
		closer, err := getMatchingCloser(char)
		if err == nil {
			s.push(closer)
			continue
		}

		if isCloser(char) {
			expected := s.pop()
			if char != expected {
				return string(char), errors.New("Expected " + string(expected) + ", but found " + string(char) + " instead")
			}
		} else {
			return "", errors.New("Corrupted input")
		}
	}
	return s.toString(), nil
}

func calcScore(ending string) int {
	score := 0
	for _, c := range ending {
		score *= 5

		switch c {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}
	return score
}

func main() {
	lines := parseInput("input.txt")
	scores := []int{}
	for _, line := range lines {
		ending, err := completeLine(line)
		if err != nil {
			fmt.Println("CORRUPTED: " + err.Error())
			continue
		}

		fmt.Println("Line is completed with", ending)
		scores = append(scores, calcScore(ending))
	}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[len(scores)/2])
}
