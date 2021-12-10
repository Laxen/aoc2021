package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type tile struct {
	value   int
	inBasin bool
}

func stringListToInt(list []string) []int {
	ret := []int{}
	for _, e := range list {
		ne, _ := strconv.Atoi(e)
		ret = append(ret, ne)
	}
	return ret
}

func parseInput(filename string) [][]tile {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)

	heightmap := [][]tile{}
	for s.Scan() {
		line := s.Text()
		heights := stringListToInt(strings.Split(line, ""))
		tiles := []tile{}
		for _, height := range heights {
			tiles = append(tiles, tile{height, false})
		}
		heightmap = append(heightmap, tiles)
	}

	return heightmap
}

func findRiskSum(heightmap [][]int) int {
	risk := 0
	for y := range heightmap {
		for x := range heightmap[0] {
			val := heightmap[y][x]

			if x-1 >= 0 {
				if heightmap[y][x-1] <= val {
					continue
				}
			}

			if x+1 < len(heightmap[0]) {
				if heightmap[y][x+1] <= val {
					continue
				}
			}

			if y-1 >= 0 {
				if heightmap[y-1][x] <= val {
					continue
				}
			}

			if y+1 < len(heightmap) {
				if heightmap[y+1][x] <= val {
					continue
				}
			}

			risk += 1 + val
		}
	}

	return risk
}

func findBasin(heightmap *[][]tile, x int, y int) int {
	if x < 0 || x >= len((*heightmap)[0]) {
		return 0
	}
	if y < 0 || y >= len((*heightmap)) {
		return 0
	}
	if (*heightmap)[y][x].inBasin {
		return 0
	}
	if (*heightmap)[y][x].value == 9 {
		return 0
	}

	sum := 1
	(*heightmap)[y][x].inBasin = true

	sum += findBasin(heightmap, x-1, y)
	sum += findBasin(heightmap, x+1, y)
	sum += findBasin(heightmap, x, y-1)
	sum += findBasin(heightmap, x, y+1)

	return sum
}

func main() {
	heightmap := parseInput("input.txt")
	sizes := []int{}
	for y := range heightmap {
		for x := range heightmap[0] {
			if heightmap[y][x].value != 9 && !heightmap[y][x].inBasin {
				size := findBasin(&heightmap, x, y)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	largest := sizes[len(sizes)-3:]
	sum := 1
	for _, bs := range largest {
		sum *= bs
	}
	fmt.Println(sum)
}
