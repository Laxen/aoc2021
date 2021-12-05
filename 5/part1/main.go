package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type board struct {
	vents [1000][1000]int
}

func (b board) show(width, height int) {
	var sb strings.Builder

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			num := b.vents[x][y]
			if num > 0 {
				fmt.Fprintf(&sb, "%d ", num)
			} else {
				fmt.Fprintf(&sb, ". ")
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func (b board) overlaps() int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if b.vents[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func minmax(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func newBoard(starts []coordinate, ends []coordinate) board {
	b := board{}

	for i := range starts {
		if starts[i].x == ends[i].x {
			x := starts[i].x
			ymin, ymax := minmax(starts[i].y, ends[i].y)

			for y := ymin; y <= ymax; y++ {
				b.vents[y][x]++
			}
		} else if starts[i].y == ends[i].y {
			xmin, xmax := minmax(starts[i].x, ends[i].x)
			y := starts[i].y

			for x := xmin; x <= xmax; x++ {
				b.vents[y][x]++
			}
		} else {
			fmt.Println("Diagonal skip!")
		}
	}

	return b
}

func parseInput(filename string) ([]coordinate, []coordinate) {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	starts := []coordinate{}
	ends := []coordinate{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")

		c1 := strings.Split(split[0], ",")
		c2 := strings.Split(split[1], ",")

		x1, _ := strconv.Atoi(c1[0])
		y1, _ := strconv.Atoi(c1[1])
		x2, _ := strconv.Atoi(c2[0])
		y2, _ := strconv.Atoi(c2[1])

		start := coordinate{
			x: x1,
			y: y1,
		}
		end := coordinate{
			x: x2,
			y: y2,
		}

		starts = append(starts, start)
		ends = append(ends, end)
	}

	return starts, ends
}

func main() {
	starts, ends := parseInput("input.txt")

	fmt.Println(starts)
	fmt.Println(ends)

	b := newBoard(starts, ends)
	b.show(10, 10)
	fmt.Println(b.overlaps())
}
