package main

import (
	. "aoc22/cube"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []Cube {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	cubeList := []Cube{}
	for s.Scan() {
		line := s.Text()
		onoff := strings.Split(line, " ")[0]
		ranges := strings.Split(strings.Split(line, " ")[1], ",")

		newCube := Cube{}
		if onoff == "on" {
			newCube.On = true
		} else {
			newCube.On = false
		}
		for i := 0; i < 3; i++ {
			s := strings.Split(ranges[i], "=")
			interval := strings.Split(s[1], "..")
			int1, _ := strconv.Atoi(interval[0])
			int2, _ := strconv.Atoi(interval[1])
			int2 += 1 // Add 1 because the cuboid should include this cube as well
			switch i {
			case 0:
				newCube.X1 = int1
				newCube.X2 = int2
			case 1:
				newCube.Y1 = int1
				newCube.Y2 = int2
			case 2:
				newCube.Z1 = int1
				newCube.Z2 = int2
			}
		}
		cubeList = append(cubeList, newCube)
	}

	return cubeList
}

func main() {
	cubes := parseInput("input.txt")

	appliedCubes := []Cube{}
	for _, newCube := range cubes {
		newAppliedCubes := []Cube{}
		for _, appliedCube := range appliedCubes {
			appliedCubeSplits := appliedCube.Subtract(newCube)
			newAppliedCubes = append(newAppliedCubes, appliedCubeSplits...)
		}

		if newCube.On {
			newAppliedCubes = append(newAppliedCubes, newCube)
		}

		appliedCubes = newAppliedCubes
	}

	sum := 0
	for _, cube := range appliedCubes {
		sum += cube.Count()
	}
	fmt.Println(sum)
}
