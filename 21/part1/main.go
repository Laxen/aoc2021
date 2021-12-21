package main

import (
	"fmt"

	m "github.com/Laxen/gohelpers/math"
)

func play(dice [100]int) ([2]int, int) {
	idx := 0
	pos := [2]int{8, 4}
	score := [2]int{}
	nRolls := 0
	for {
		for p := range pos {
			sum := dice[idx%100] + dice[(idx+1)%100] + dice[(idx+2)%100]
			idx = (idx + 3) % 100
			nRolls += 3
			pos[p] = (pos[p] + sum) % 10
			if pos[p] == 0 {
				pos[p] = 10
			}
			score[p] += pos[p]
			if score[p] >= 1000 {
				return score, nRolls
			}
		}
	}
}

func main() {
	dice := [100]int{}
	for i := 0; i < 100; i++ {
		dice[i] = i + 1
	}

	score, nRolls := play(dice)
	fmt.Println(score)
	fmt.Println(nRolls)

	fmt.Println(m.Min(score[0], score[1]) * nRolls)
}
