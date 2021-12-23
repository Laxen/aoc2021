package main

import "fmt"

type State struct {
	PositionP1 int
	PositionP2 int
	ScoreP1    int
	ScoreP2    int
}

func getWins(state State, stepMap map[int]int) (winsP1, winsP2 uint64) {
	for step, nUnivsP1 := range stepMap {
		stateP1 := state

		stateP1.PositionP1 += step
		if stateP1.PositionP1 > 10 {
			stateP1.PositionP1 -= 10
		}
		stateP1.ScoreP1 += stateP1.PositionP1

		if stateP1.ScoreP1 >= 21 {
			winsP1 += uint64(nUnivsP1)
			continue
		}

		for step, nUnivsP2 := range stepMap {
			stateP2 := stateP1
			stateP2.PositionP2 += step
			if stateP2.PositionP2 > 10 {
				stateP2.PositionP2 -= 10
			}
			stateP2.ScoreP2 += stateP2.PositionP2

			if stateP2.ScoreP2 >= 21 {
				winsP2 += uint64(nUnivsP1 * nUnivsP2)
				continue
			}

			wP1, wP2 := getWins(stateP2, stepMap)
			winsP1 += wP1 * uint64(nUnivsP1*nUnivsP2)
			winsP2 += wP2 * uint64(nUnivsP1*nUnivsP2)
		}
	}

	return winsP1, winsP2
}

func main() {
	state := State{
		PositionP1: 8,
		PositionP2: 4,
		ScoreP1:    0,
		ScoreP2:    0,
	}

	stepMap := map[int]int{ // map[stepIncrease]nUniverses
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	winsP1, winsP2 := getWins(state, stepMap)
	fmt.Println(winsP1, winsP2)
}
