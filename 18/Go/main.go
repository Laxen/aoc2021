package main

import (
	"fmt"
	"regexp"
	"time"
)

func findLeaves(snailNumber string) [][]int {
	re := regexp.MustCompile(`\[\d,\d\]`)
	matches := re.FindAllStringIndex(snailNumber, -1)
	return matches
}

func findParent(snailNumber string, indices []int) {
	commaIdx := func() int {
		idx := indices[0] - 1
		if idx >= 0 {
			if snailNumber[idx] == ',' {
				return idx
			}
		}
		return indices[1]
	}()

	fmt.Println(commaIdx)
}

func parseSnailNumber(sn string) {
	leaves := findLeaves(sn)
	for _, leaf := range leaves {
		go findParent(sn, leaf)
	}
}

func main() {
	// sn := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"
	sn := "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"

	parseSnailNumber(sn)

	time.Sleep(time.Second)
}
