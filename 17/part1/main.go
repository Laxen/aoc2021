package main

import (
	"fmt"
	"math"
)

func calcVelX(goalX int, startX int, steps int) float32 {
	return float32(goalX-startX+(steps-1)*steps/2) / float32(steps)
}

func calcStepsForStaticX(xEnd int) int {
	n1 := -0.5 + math.Sqrt(0.25+2*float64(xEnd))

	return int(n1 + 0.5)
}

func calcVy(nMin int, nMax int, yEndMin int, yEndMax int) ([]int, int) {
	vys := []int{}
	yMax := -1000000000

	for n := nMin; n < nMin+1000; n++ {
		for vy := 0; vy < 10000; vy++ {
			yn := n*vy - (n-1)*n/2
			if yn >= yEndMin && yn <= yEndMax {
				vys = append(vys, vy)
			}
		}
	}

	for _, vy := range vys {
		for n := 0; n < 1000; n++ {
			y := n*vy - (n-1)*n/2
			if y > yMax {
				yMax = y
			}
		}
	}

	return vys, yMax
}

func main() {
	nMin := calcStepsForStaticX(20)
	nMax := calcStepsForStaticX(30)
	vy, yMax := calcVy(nMin, nMax, -10, -5)
	// nMin := calcStepsForStaticX(56)
	// nMax := calcStepsForStaticX(76)
	// vy, yMax := calcVy(nMin, nMax, -162, -134)
	fmt.Println(vy, yMax)
}
