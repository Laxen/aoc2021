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

func calcVy(nMin int, nMax int, yEndMin int, yEndMax int) map[int][]int {
	vysPerStep := map[int][]int{}

	for n := nMin; n < nMin+1000; n++ {
		vys := []int{}
		for vy := 0; vy < 10000; vy++ {
			yn := n*vy - (n-1)*n/2
			if yn >= yEndMin && yn <= yEndMax {
				vys = append(vys, vy)
			}
		}
		vysPerStep[n] = vys
	}

	return vysPerStep
}

func calcXPos(steps int, vx int) int {
	x := 0
	for n := 0; n < steps; n++ {
		x += vx
		if vx > 0 {
			vx--
		}
	}
	return x
}

func calcPossibleVelsForStep(n int, xMin int, xMax int, yMin int, yMax int) ([]int, []int) {
	vxs := []int{}
	vys := []int{}

	for vx := 1; vx <= xMax; vx++ {
		// x := n*vx - n*(n-1)/2
		x := calcXPos(n, vx)
		/* TODO: This needs a cap on vx, can't go below 0 */
		/* You already have a function for max n, use that? */
		if x >= xMin && x <= xMax {
			vxs = append(vxs, vx)
		}
	}

	for vy := yMin; vy < 10000; vy++ {
		y := n*vy - n*(n-1)/2
		if y >= yMin && y <= yMax {
			vys = append(vys, vy)
		}
	}

	return vxs, vys
}

func isElementInList(list []int, e int) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}
	return false
}

func removeDuplicates(m map[int][]int) {
	for k, v := range m {
		newV := []int{}
		for _, e := range v {
			if !isElementInList(newV, e) {
				newV = append(newV, e)
			}
		}
		m[k] = newV
	}
}

func main() {
	// nMin := calcStepsForStaticX(20)
	// nMax := calcStepsForStaticX(30)
	// vysPerStep := calcVy(nMin, nMax, -10, -5)

	// nMin := calcStepsForStaticX(56)
	// nMax := calcStepsForStaticX(76)
	// vy, yMax := calcVy(nMin, nMax, -162, -134)

	// xMin := 20
	// xMax := 30
	// yMin := -10
	// yMax := -5
	xMin := 56
	xMax := 76
	yMin := -162
	yMax := -134

	vels := map[int][]int{}
	for n := 1; n <= 10000; n++ {
		vxs, vys := calcPossibleVelsForStep(n, xMin, xMax, yMin, yMax)
		for _, vx := range vxs {
			if vxmap, ok := vels[vx]; ok {
				vxmap = append(vxmap, vys...)
				vels[vx] = vxmap
			} else {
				vels[vx] = vys
			}
		}
	}
	removeDuplicates(vels)
	count := 0
	for _, vys := range vels {
		count += len(vys)
	}
	fmt.Println(vels)
	fmt.Println(count)

	// 975 too low
}
