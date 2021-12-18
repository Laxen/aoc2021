package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	val int
	a   *pair
	b   *pair
}

func (p pair) toString() string {
	if p.val >= 0 {
		return strconv.Itoa(p.val)
	} else {
		return fmt.Sprintf("[%s,%s]", p.a.toString(), p.b.toString())
	}
}

func addPairs(l pair, r pair) pair {
	return pair{-1, &l, &r}
}

func parsePair(a string, b string) pair {

}

func parseInput(filename string) *pair {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	root := pair{}
	for s.Scan() {
		line := s.Text()
		line = line[1 : len(line)-2]
		vals := strings.Split(line, ",")
		a, _ := strconv.Atoi(vals[0])
		b, _ := strconv.Atoi(vals[1])
		apair := pair{a, nil, nil}
		bpair := pair{b, nil, nil}
		root = pair{-1, &apair, &bpair}
	}

	return &root
}

func main() {
	pair := parseInput("example.txt")
	fmt.Println(pair.toString())
}
