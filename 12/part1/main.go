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

func findPaths(caveMap map[string][]string, cave string, path []string, paths [][]string) [][]string {
	if countStrings(path, cave) > 0 && isLowerCase(cave) {
		return paths
	} else if cave == "end" {
		path = append(path, "end")
		paths = append(paths, path)
		return paths
	}

	path = append(path, cave)
	for _, neighbor := range caveMap[cave] {
		paths = findPaths(caveMap, neighbor, path, paths)
	}

	return paths
}

func main() {
	caveMap := parseInput("input.txt")

	paths := findPaths(caveMap, "start", []string{}, [][]string{})
	fmt.Println(len(paths))
}
