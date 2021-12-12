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

type cave struct {
	name      string
	small     bool
	visited   bool
	neighbors []string
}

func makeCave(name string) (c cave) {
	c.name = name
	c.visited = false

	c.small = true
	for _, r := range name {
		if unicode.IsUpper(r) {
			c.small = false
			break
		}
	}

	return c
}

func (c cave) addNeighbor(name string) cave {
	c.neighbors = append(c.neighbors, name)
	return c
}

func parseInput(filename string) map[string]cave {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	caveMap := map[string]cave{}
	for s.Scan() {
		line := s.Text()
		nodes := strings.Split(line, "-")

		if _, ok := caveMap[nodes[0]]; !ok {
			caveMap[nodes[0]] = makeCave(nodes[0])
		}
		caveMap[nodes[0]] = caveMap[nodes[0]].addNeighbor(nodes[1])

		if _, ok := caveMap[nodes[1]]; !ok {
			caveMap[nodes[1]] = makeCave(nodes[1])
		}
		caveMap[nodes[1]] = caveMap[nodes[1]].addNeighbor(nodes[0])
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

func findPaths(caveMap map[string]cave, curCaveName string, path []string, paths [][]string) [][]string {
	curCave := caveMap[curCaveName]

	if countStrings(path, curCaveName) > 0 && curCave.small {
		return paths
	} else if curCave.name == "end" {
		path = append(path, "end")
		paths = append(paths, path)
		return paths
	}

	path = append(path, curCaveName)
	for _, neighbor := range curCave.neighbors {
		paths = findPaths(caveMap, neighbor, path, paths)
	}

	return paths
}

func main() {
	caveMap := parseInput("input.txt")

	paths := findPaths(caveMap, "start", []string{}, [][]string{})
	fmt.Println(len(paths))
}
