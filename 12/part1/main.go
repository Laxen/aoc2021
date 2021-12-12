// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseInput(filename string) map[string][]string {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	caveMap := map[string][]string{}
	for s.Scan() {
		line := s.Text()
		nodes := strings.Split(line, "-")

		c0 := []string{}
		c1 := []string{}
		if _, ok := caveMap[nodes[0]]; ok {
			c0 = caveMap[nodes[0]]
		}
		if _, ok := caveMap[nodes[1]]; ok {
			c1 = caveMap[nodes[1]]
		}
		c0 = append(c0, nodes[1])
		c1 = append(c1, nodes[0])
		caveMap[nodes[0]] = c0
		caveMap[nodes[1]] = c1
	}

	return caveMap
}

func countStrings(list []string, target string) (ret int) {
	for _, s := range list {
		if s == target {
			ret++
		}
	}
	return ret
}

func isLowerCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func findPaths(caveMap map[string][]string, cave string, caveCount map[string]int, paths int) int {
	if caveCount[cave] > 0 && isLowerCase(cave) {
		if cave == "start" {
			return paths
		}
		for c, cnt := range caveCount {
			if isLowerCase(c) && cnt > 1 {
				return paths
			}
		}
	} else if cave == "end" {
		return paths + 1
	}

	caveCount[cave]++
	for _, neighbor := range caveMap[cave] {
		paths = findPaths(caveMap, neighbor, caveCount, paths)
	}

	caveCount[cave]--
	return paths
}

func main() {
	caveMap := parseInput("input.txt")

	paths := findPaths(caveMap, "start", map[string]int{}, 0)
	fmt.Println(paths)
}
