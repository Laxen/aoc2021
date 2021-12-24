package main

import (
	"strconv"
	"strings"

	m "github.com/Laxen/gohelpers/math"
)

func executeInstruction(instruction string, memory map[string]int, input []int) (map[string]int, []int) {
	instructions := strings.Split(instruction, " ")
	switch instructions[0] {
	case "inp":
		data := input[0]
		input = input[1:]
		memory[instructions[1]] = data
	case "add":
		i, err := strconv.Atoi(instructions[2])
		if err != nil {
			i = memory[instructions[2]]
		}
		memory[instructions[1]] += i
	case "mul":
		i, err := strconv.Atoi(instructions[2])
		if err != nil {
			i = memory[instructions[2]]
		}
		memory[instructions[1]] *= i
	case "div":
		i, err := strconv.Atoi(instructions[2])
		if err != nil {
			i = memory[instructions[2]]
		}
		memory[instructions[1]] /= i
	case "mod":
		i, err := strconv.Atoi(instructions[2])
		if err != nil {
			i = memory[instructions[2]]
		}
		memory[instructions[1]] %= i
	case "eql":
		i, err := strconv.Atoi(instructions[2])
		if err != nil {
			i = memory[instructions[2]]
		}

		if memory[instructions[1]] == i {
			memory[instructions[1]] = 1
		} else {
			memory[instructions[1]] = 0
		}
	}

	return memory, input
}

func execute(instructions []string, memory map[string]int, input []int) map[string]int {
	for _, instruction := range instructions {
		memory, input = executeInstruction(instruction, memory, input)
	}
	return memory
}

func speedyExecute(serno int) int {
	divs := []int{
		1, 1, 1, 26, 1, 1, 1, 26, 26, 1, 26, 26, 26, 26,
	}

	adds1 := []int{
		13, 11, 14, -5, 14, 10, 12, -14, -8, 13, 0, -5, -9, -1,
	}

	adds2 := []int{
		0, 3, 8, 5, 13, 9, 6, 1, 1, 2, 7, 5, 8, 15,
	}

	z := 0
	for di := 0; di < 14; di++ {
		digit := (serno / int(m.Pow(10, 13-di))) % 10
		if digit == 0 {
			return 1
		}

		/*
			divs has 7 1s and 7 26s. Every time divs is 1 it will add a number
			>10 (adds1) to digit so x will ALWAYS be 1. This means that z will
			ALWAYS grow with a factor of 26 + digit + adds2. In order to get z=0
			again we have to reach z=0 using the 7 26 divisions. In order to do
			that z has to ALWAYS be reduced by a factor of 26 when divs == 26.
			That means that x has to be 0 every time divs == 26. If that's not
			the case then the serial number is invalid.

			x := z%26 + adds1[di]
			This in turn means that x == digit gives z%26 - adds1[di] == digit
			z%26 is always the PREVIOUS digit+adds2[di], so we get this
			constraint every time divs is 26:

			prevDigit + adds2[prevDi] + adds1[di] == digit

			This can be deconstructed into the following constraints
			serno[2] + 8 - 5 == serno[3]
			serno[6] + 6 - 14 == serno[7]
			serno[7] + 1 - 8 == serno[8]
			serno[9] + 2 - 0 == serno[10]
			serno[10] + 7 - 5 == serno[11]
			serno[11] + 5 - 9 == serno[12]
			serno[12] + 8 - 1 == serno[13]

			-----------

			w = digit
			x = z%26+adds1[di]
			z /= divs[di]

			x = 1 if x != w else 0

			z = z * (25*x + 1) + (w + adds2[di]) * x

			---------------

			if z%26 + adds1[di] != digit
				z = z*26/divs[di] + digit + adds2[di]
			else
				z = z/divs[di]
		*/

		x := z%26 + adds1[di]
		if x == digit {
			x = 0
		} else {
			x = 1
		}

		z /= divs[di]
		z = z*26 + digit + adds2[di]
	}

	return z
}

func main() {
	// instructions := []string{
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 13",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 0",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 11",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 3",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 14",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 8",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -5",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 5",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 14",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 13",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 10",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 9",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 12",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 6",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -14",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 1",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -8",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 1",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 1",
	// 	"add x 13",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 2",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x 0",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 7",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -5",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 5",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -9",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 8",
	// 	"mul y x",
	// 	"add z y",
	// 	"inp w",
	// 	"mul x 0",
	// 	"add x z",
	// 	"mod x 26",
	// 	"div z 26",
	// 	"add x -1",
	// 	"eql x w",
	// 	"eql x 0",
	// 	"mul y 0",
	// 	"add y 25",
	// 	"mul y x",
	// 	"add y 1",
	// 	"mul z y",
	// 	"mul y 0",
	// 	"add y w",
	// 	"add y 15",
	// 	"mul y x",
	// 	"add z y",
	// }

	// memory := map[string]int{}
	// memory = execute(instructions, memory, []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9})
	// z := speedyExecute(13579246899999)
	// fmt.Println(memory)
	// fmt.Println(z)

	// panic("ok")

	// for serno := 99999999999999; serno > 11111111111111; serno-- {
	// 	if serno%10000000 == 0 {
	// 		fmt.Println(serno)
	// 	}

	// 	z := speedyExecute(serno)
	// 	if z == 0 {
	// 		fmt.Printf("%d is valid\n", serno)
	// 	}
	// }

	/*
		serno[2] + 8 - 5 == serno[3]
		serno[6] + 6 - 14 == serno[7]
		serno[7] + 1 - 8 == serno[8]
		serno[9] + 2 - 0 == serno[10]
		serno[10] + 7 - 5 == serno[11]
		serno[11] + 5 - 9 == serno[12]
		serno[12] + 8 - 1 == serno[13]

		serno[2]  + 3 == serno[3]
		serno[6]  - 8 == serno[7]
		serno[7]  - 7 == serno[8]
		serno[9]  + 2 == serno[10]
		serno[10] + 2 == serno[11]
		serno[11] - 4 == serno[12]
		serno[12] + 7 == serno[13]

		serno[7]  = serno[6] - 6
		serno[8]  = serno[6] - 15 ??????????????????
		serno[9]  = whatever
		serno[10] = serno[9] + 2
		serno[11] = serno[9] + 4
		serno[12] = serno[9]
		serno[13] = serno[9] + 7
	*/

	for 
}
