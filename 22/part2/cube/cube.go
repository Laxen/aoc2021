package cube

import "fmt"

type Cube struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
	On     bool
}

func (c Cube) IsEmpty() bool {
	if c.X1 == c.X2 || c.Y1 == c.Y2 || c.Z1 == c.Z2 {
		return true
	}
	return false
}

func (c Cube) Print() {
	s := "off"
	if c.On {
		s = "on"
	}
	fmt.Printf("%s: x=%d..%d,y=%d..%d,z=%d..%d\n", s, c.X1, c.X2, c.Y1, c.Y2, c.Z1, c.Z2)
}

func (a Cube) Count() int {
	return (a.X2 - a.X1) * (a.Y2 - a.Y1) * (a.Z2 - a.Z1)
}

func (a Cube) intersection1D(b Cube, dimension string) Cube {
	ret := Cube{}
	ret.On = a.On

	aD1, aD2, bD1, bD2 := 0, 0, 0, 0
	retD1, retD2 := 0, 0

	switch dimension {
	case "x":
		aD1 = a.X1
		aD2 = a.X2
		bD1 = b.X1
		bD2 = b.X2
	case "y":
		aD1 = a.Y1
		aD2 = a.Y2
		bD1 = b.Y1
		bD2 = b.Y2
	case "z":
		aD1 = a.Z1
		aD2 = a.Z2
		bD1 = b.Z1
		bD2 = b.Z2
	}

	if bD1 >= aD1 {
		if bD1 >= aD2 {
			return ret // No intersection, case 3
		}
		if bD2 >= aD2 { // Case 1
			retD1 = bD1
			retD2 = aD2
		} else { // Case 4
			retD1 = bD1
			retD2 = bD2
		}
	} else {
		if aD1 >= bD2 { // No intersection, case 6
			return ret
		}
		if aD2 >= bD2 { // Case 2
			retD1 = aD1
			retD2 = bD2
		} else { // Case 5
			retD1 = aD1
			retD2 = aD2
		}
	}

	switch dimension {
	case "x":
		ret.X1 = retD1
		ret.X2 = retD2
	case "y":
		ret.Y1 = retD1
		ret.Y2 = retD2
	case "z":
		ret.Z1 = retD1
		ret.Z2 = retD2
	}

	return ret
}

func (a Cube) Intersection(b Cube) Cube {
	ret := Cube{}

	ret.On = a.On

	intersectionX := a.intersection1D(b, "x")
	intersectionY := a.intersection1D(b, "y")
	intersectionZ := a.intersection1D(b, "z")

	ret.X1 = intersectionX.X1
	ret.X2 = intersectionX.X2
	ret.Y1 = intersectionY.Y1
	ret.Y2 = intersectionY.Y2
	ret.Z1 = intersectionZ.Z1
	ret.Z2 = intersectionZ.Z2

	if ret.IsEmpty() {
		return Cube{}
	}

	return ret
}

func subtract1D(x1, x2, y1, y2 int) []Cube {
	ret := []Cube{}

	c1 := Cube{}
	c1.X1 = x1
	c1.X2 = y1

	c2 := Cube{}
	c2.X1 = y2
	c2.X2 = x2

	if c1.X1 != c1.X2 {
		ret = append(ret, c1)
	}
	if c2.X1 != c2.X2 {
		ret = append(ret, c2)
	}

	return ret
}

func (a Cube) Subtract(b Cube) []Cube {
	ret := []Cube{}

	intersection := a.Intersection(b)
	if intersection.IsEmpty() {
		ret = append(ret, a)
		return ret
	}

	subX := subtract1D(a.X1, a.X2, intersection.X1, intersection.X2)
	for _, c := range subX {
		c.On = a.On
		c.Y1 = a.Y1
		c.Y2 = a.Y2
		c.Z1 = a.Z1
		c.Z2 = a.Z2
		ret = append(ret, c)
	}

	subY := subtract1D(a.Y1, a.Y2, intersection.Y1, intersection.Y2)
	for _, c := range subY {
		c.Y1 = c.X1
		c.Y2 = c.X2

		c.On = a.On
		c.X1 = intersection.X1
		c.X2 = intersection.X2
		c.Z1 = a.Z1
		c.Z2 = a.Z2
		ret = append(ret, c)
	}

	subZ := subtract1D(a.Z1, a.Z2, intersection.Z1, intersection.Z2)
	for _, c := range subZ {
		c.Z1 = c.X1
		c.Z2 = c.X2

		c.On = a.On
		c.X1 = intersection.X1
		c.X2 = intersection.X2
		c.Y1 = intersection.Y1
		c.Y2 = intersection.Y2
		ret = append(ret, c)
	}

	return ret
}
