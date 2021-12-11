package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type octopus struct {
	energy     int
	hasFlashed bool
}

func parseInput(filename string) [][]octopus {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	octomap := [][]octopus{}
	for s.Scan() {
		line := s.Text()
		energies := strings.Split(line, "")
		octoline := []octopus{}
		for _, energy := range energies {
			energy, _ := strconv.Atoi(energy)
			octo := octopus{energy, false}
			octoline = append(octoline, octo)
		}
		octomap = append(octomap, octoline)
	}

	return octomap
}

func calcFlash(octos *[][]octopus, x, y int) {
	if x < 0 || y < 0 || x > len((*octos)[0])-1 || y > len(*octos)-1 {
		return
	}

	octo := (*octos)[y][x]
	octo.energy++

	if octo.energy <= 9 || octo.hasFlashed {
		(*octos)[y][x] = octo
		return
	}

	octo.hasFlashed = true
	(*octos)[y][x] = octo

	calcFlash(octos, x-1, y)
	calcFlash(octos, x+1, y)
	calcFlash(octos, x, y-1)
	calcFlash(octos, x, y+1)
	calcFlash(octos, x+1, y+1)
	calcFlash(octos, x-1, y+1)
	calcFlash(octos, x+1, y-1)
	calcFlash(octos, x-1, y-1)
}

func step(octos *[][]octopus) int {
	for y, octoline := range *octos {
		for x := range octoline {
			calcFlash(octos, x, y)
		}
	}

	flashes := 0
	for y, octoline := range *octos {
		for x, octo := range octoline {
			if octo.hasFlashed {
				(*octos)[y][x].hasFlashed = false
				(*octos)[y][x].energy = 0
				flashes++
				fmt.Printf("[%d]\t", 0)
			} else {
				fmt.Printf("%d\t", octo.energy)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

	return flashes
}

func main() {
	octos := parseInput("input.txt")

	for i := 0; i < 1000; i++ {
		flashes := step(&octos)
		if flashes == 100 {
			fmt.Println(i + 1)
			break
		}
	}
}
