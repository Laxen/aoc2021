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

func findPaths(caveMap map[string]cave, curCaveName string, curPath []cave, paths [][]cave, doublevisit bool) [][]cave {
	curCave := caveMap[curCaveName]

	if curCave.visited && curCave.small {
		if doublevisit || curCaveName == "start" {
			// fmt.Printf("[%s] Already visited and small, returning...\n", curCaveName)
			return paths
		}
		doublevisit = true
	} else if curCave.name == "end" {
		curPath = append(curPath, curCave)
		/* Need to copy the slice otherwise it will be overwritten due to memory sharing... */
		newPath := make([]cave, len(curPath))
		copy(newPath, curPath)
		paths = append(paths, newPath)

		// fmt.Printf("[%s] This is the end, curPath is %v\n", curCaveName, curPath)
		// fmt.Printf("[%s] Paths are:\n", curCaveName)
		// for _, path := range paths {
		// 	fmt.Printf("[%s] %v\n", curCaveName, path)
		// }
		return paths
	}

	curCave.visited = true
	curPath = append(curPath, curCave)
	// fmt.Printf("[%s] Marked as visited\n", curCaveName)
	// fmt.Printf("[%s] Current path is %v\n", curCaveName, curPath)

	newCaveMap := map[string]cave{}
	for k, c := range caveMap {
		newCaveMap[k] = c
	}
	newCaveMap[curCaveName] = curCave
	for _, neighbor := range curCave.neighbors {
		// fmt.Printf("[%s] Go to %s\n", curCaveName, neighbor)
		paths = findPaths(newCaveMap, neighbor, curPath, paths, doublevisit)
		// fmt.Printf("[%s] Returned from %s\n", curCaveName, neighbor)
	}

	// fmt.Printf("[%s] No more paths, returning...\n", curCaveName)
	return paths
}

func test(caveMap map[string]cave, curCaveName string) {
	curCave := caveMap[curCaveName]
	curCave.visited = true
	caveMap[curCaveName] = curCave
}

func main() {
	caveMap := parseInput("input.txt")

	paths := findPaths(caveMap, "start", []cave{}, [][]cave{}, false)
	// for _, path := range paths {
	// 	for _, c := range path {
	// 		fmt.Printf("%s ", c.name)
	// 	}
	// 	fmt.Printf("\n")
	// }
	fmt.Println(len(paths))
}
