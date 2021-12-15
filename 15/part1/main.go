package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type node struct {
	x, y    int
	risk    int
	f, g, h float64
	parent  *node
}

func parseInput(filename string) [][]node {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	nodes := [][]node{}
	y := 0
	for s.Scan() {
		line := s.Text()
		nodeLine := []node{}
		for x, c := range line {
			cint, _ := strconv.Atoi(string(c))
			n := node{x, y, cint, -1, -1, -1, nil}
			nodeLine = append(nodeLine, n)
		}
		nodes = append(nodes, nodeLine)
		y++
	}

	nodes[0][0].f = 0
	nodes[0][0].g = 0
	nodes[0][0].h = 0
	return nodes
}

func isNodeInList(n node, l []*node) bool {
	for _, e := range l {
		if n == *e {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func dist(n1 node, n2 node) float64 {
	dx := float64(abs(n1.x - n2.x))
	dy := float64(abs(n1.y - n2.y))

	ret := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	return float64(ret)
}

func astar(nodes [][]node) *node {
	open := []*node{}
	close := []*node{}

	open = append(open, &nodes[0][0])

	for len(open) > 0 {
		currentNode := open[0]
		currentNodeIndex := 0
		for i, node := range open {
			if node.f < currentNode.f {
				currentNode = node
				currentNodeIndex = i
			}
		}

		// fmt.Printf("Checking node (%d,%d) G=%f\n", currentNode.x, currentNode.y, currentNode.g)

		open[currentNodeIndex] = open[len(open)-1]
		open = open[:len(open)-1]

		close = append(close, currentNode)

		if currentNode.y == len(nodes)-1 && currentNode.x == len(nodes[0])-1 {
			return currentNode
		}

		children := []*node{}
		if currentNode.x > 0 {
			children = append(children, &nodes[currentNode.y][currentNode.x-1])
		}
		if currentNode.x < len(nodes[0])-1 {
			children = append(children, &nodes[currentNode.y][currentNode.x+1])
		}
		if currentNode.y > 0 {
			children = append(children, &nodes[currentNode.y-1][currentNode.x])
		}
		if currentNode.y < len(nodes)-1 {
			children = append(children, &nodes[currentNode.y+1][currentNode.x])
		}

		for _, child := range children {
			if isNodeInList(*child, close) {
				continue
			}

			childG := currentNode.g + float64(child.risk)
			childH := dist(*child, nodes[len(nodes)-1][len(nodes[0])-1])
			childF := childG + childH

			if isNodeInList(*child, open) {
				if child.g != -1 && childG > child.g {
					continue
				}
			} else {
				open = append(open, child)
			}

			child.g = childG
			child.h = childH
			child.f = childF
			child.parent = currentNode
			// fmt.Printf("    Child(%d,%d) G=%f, H=%f, F=%f\n", child.x, child.y, child.g, child.h, child.f)
		}
	}

	return nil
}

func test(nodes [][]node) {
	open := []*node{}
	open = append(open, &nodes[0][0])
	currentNode := open[0]
	currentNode.x = 1000
}

func main() {
	nodes := parseInput("input.txt")

	// fmt.Println(nodes[0][0])
	// test(nodes)
	// fmt.Println(nodes[0][0])
	endNode := astar(nodes)
	n := endNode
	risk := 0
	for n != nil {
		fmt.Println(n.x, n.y)
		risk += n.risk
		n = n.parent
	}
	fmt.Println(risk - nodes[0][0].risk)
}
