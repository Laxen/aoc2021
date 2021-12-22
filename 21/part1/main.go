package main

import (
	"fmt"
	"sort"
)

func update(player map[int]map[int]uint64, totalUnivs uint64, stepMap map[int]uint64) (map[int]map[int]uint64, uint64) {
	var playerScore uint64
	newPlayer := map[int]map[int]uint64{}
	for i := 0; i < 10; i++ {
		newPlayer[i] = map[int]uint64{}
	}

	for tile := 0; tile < 10; tile++ {
		for step, nNewUnivs := range stepMap {
			newTile := (tile + step) % 10
			for score, nUnivs := range player[tile] {
				newScore := score + newTile + 1
				addUnivs := nNewUnivs * nUnivs
				// if totalUnivs > 0 {
				// 	addUnivs *= 27
				// }
				if newScore <= 21 {
					newPlayer[newTile][newScore] += addUnivs
				} else {
					playerScore += (totalUnivs) * 27
				}
			}
		}
	}

	return newPlayer, playerScore
}

func play2() {
	// scoreMaps := [2]map[int]int{} // map[score]nUniverses

	stepMap := map[int]uint64{ // map[stepIncrease]nUniverses
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	// posMap := map[int]int{ // map[pos]nUniverses
	// 	4: 1,
	// }

	// scoreMap := map[int]int{} // map[pos]maxScore

	// map[pos]map[score]nUniverses
	// Every position has a map that shows how many universes have a certain score on that position
	// For every round the player goes through their positions
	//   For every position the player has, they go through every score
	//     For every score the player has on that position, they take all universes that have that score, multiply it by all universes the other player has in total,

	// map[tile]map[score]nUnivs
	p1 := map[int]map[int]uint64{}
	p2 := map[int]map[int]uint64{}
	for i := 0; i < 10; i++ {
		p1[i] = map[int]uint64{}
		p2[i] = map[int]uint64{}
	}

	p1[3][0] = 1
	p2[7][0] = 1

	var p1Score uint64
	var p2Score uint64

	for rounds := 0; rounds < 3; rounds++ {
		fmt.Println("ROUND", rounds)
		fmt.Println()

		fmt.Println("Number of univs:", numberOfUnivs(p2))
		newP1, p1ScoreIncrease := update(p1, numberOfUnivs(p2), stepMap)
		// newP1, p1ScoreIncrease := update(p1, 2*uint64(rounds), stepMap)
		p1 = newP1
		p1Score += p1ScoreIncrease

		fmt.Println("P1 score:", p1Score)
		printMap(p1)
		fmt.Println()

		fmt.Println("Number of univs:", numberOfUnivs(p1))
		newP2, p2ScoreIncrease := update(p2, numberOfUnivs(p1), stepMap)
		// newP2, p2ScoreIncrease := update(p2, 2*uint64(rounds)+1, stepMap)
		p2 = newP2
		p2Score += p2ScoreIncrease

		fmt.Println("P2 score:", p2Score)
		printMap(p2)
		fmt.Println()
	}

	// for i := 0; i < 3; i++ {
	// 	newPosMap := map[int]int{}
	// 	newScoreMap := map[int]int{}
	// 	for s, nUniverses := range stepMap {
	// 		for p := range posMap {
	// 			newPos := (p + s) % 10
	// 			if newPos == 0 {
	// 				newPos = 10
	// 			}
	// 			newPosMap[newPos] += posMap[p] * nUniverses
	// 			newScoreMap[newPos] = m.Max(newScoreMap[newPos], scoreMap[p]+newPos)
	// 		}
	// 	}
	// 	posMap = newPosMap
	// 	scoreMap = newScoreMap
	// 	fmt.Println(posMap)
	// 	fmt.Println(scoreMap)
	// 	fmt.Println()
	// }
}

func printMap(p map[int]map[int]uint64) {
	keys := make([]int, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, p[k])
	}
}

func numberOfUnivs(p map[int]map[int]uint64) (n uint64) {
	for _, scoreMap := range p {
		for _, nUnivs := range scoreMap {
			n += nUnivs
		}
	}
	return n
}

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
	play2()
}
