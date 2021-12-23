package main

import (
	"fmt"

	m "github.com/Laxen/gohelpers/math"
)

type Pod struct {
	id         string
	hallwayPos int
	roomPos    int
	state      int
}

func (p Pod) Cost() int {
	switch p.id {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	}
	panic("Pod has invalid ID")
}

func (p Pod) AllowedToEnterRoom(idx int) bool {
	switch p.id {
	case "A":
		return idx == 2
	case "B":
		return idx == 4
	case "C":
		return idx == 6
	case "D":
		return idx == 8
	}
	panic("Pod has invalid ID")
}

func validPath(pods []Pod, pod Pod, start, end int) bool {
	if end == 2 || end == 4 || end == 6 || end == 8 {
		if !pod.AllowedToEnterRoom(end) {
			return false
		}
	}

	s := m.Min(start, end)
	e := m.Max(start, end)

	for i := s; i < e; i++ {
		for _, p := range pods {
			if p.hallwayPos == i {
				// Check if pod is in the hallway, in that case it's blocking
				if p.roomPos == -1 {
					return false
				}
				// Check if pod is in the bottom of a room and is allowed to be
				// there, then we can put this pod on top of it
				if p.roomPos == 0 && p.AllowedToEnterRoom(i) {
					return true
				}
			}
		}
	}

	// No pod is blocking the hallway. If e is a room spot then check if the room is available.
	if e == 2 || e == 4 || e == 6 || e == 8 {
		for _, p := range pods {
			if p.hallwayPos == e {
				// Check if pod is in the bottom of a room and is allowed to be
				// there, then we can put this pod on top of it. Otherwise
				// there's no valid path.
				if p.roomPos == 0 && p.AllowedToEnterRoom(e) {
					return true
				}
				return false
			}
		}
	}

	return true
}

func printMap(hallway [11]Pod, rooms []Room) {
	for _, pod := range hallway {
		switch pod.cost {
		case 0:
			fmt.Printf(".")
		case 1:
			fmt.Printf("A")
		case 10:
			fmt.Printf("B")
		case 100:
			fmt.Printf("C")
		case 1000:
			fmt.Printf("D")
		}
	}
	for i := 0; i < 2; i++ {
		fmt.Println()
		fmt.Printf("  ")
		for _, room := range rooms {
			switch room.pods[i].cost {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf("A")
			case 10:
				fmt.Printf("B")
			case 100:
				fmt.Printf("C")
			case 1000:
				fmt.Printf("D")
			}
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	hallway := [11]Pod{}
	rooms := []Room{
		Room{
			2,
			[]Pod{
				Pod{10},
				Pod{100},
			},
		},
		Room{
			4,
			[]Pod{
				Pod{10},
				Pod{100},
			},
		},
		Room{
			6,
			[]Pod{
				Pod{1000},
				Pod{1},
			},
		},
		Room{
			8,
			[]Pod{
				Pod{1000},
				Pod{1},
			},
		},
	}

	pathfind(hallway, rooms)
}
