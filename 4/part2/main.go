package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type number struct {
    value int
    marked bool
}

type board struct {
    numbers [5][5] number
    hasWon bool
}

func (b board) show() {
    var sb strings.Builder
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if b.numbers[x][y].marked {
                fmt.Fprintf(&sb, "*%s*\t", strconv.Itoa(b.numbers[x][y].value))
            } else {
                fmt.Fprintf(&sb, "%s\t", strconv.Itoa(b.numbers[x][y].value))
            }
        }
        sb.WriteString("\n")
    }

    fmt.Println(sb.String())
}

func (b *board) mark(number int) bool {
    lineCount := 0
    colCount := [5]int{}

    for x := 0; x < 5; x++ {
        lineCount = 0
        for y := 0; y < 5; y++ {
            n := &b.numbers[x][y]
            if n.value == number {
                n.marked = true
            }

            if n.marked {
                lineCount++
                colCount[y]++

                if colCount[y] == 5 {
                    return true
                }
            }
        }

        if lineCount == 5 {
            return true
        }
    }

    return false
}

func (b board) hasBingo() bool {
    lineCount := 0
    var colCount [5]int

    for x := 0; x < 5; x++ {
        lineCount = 0
        for y := 0; y < 5; y++ {
            if b.numbers[x][y].marked {
                lineCount++
                colCount[y]++

                if colCount[y] == 5 {
                    return true
                }
            }
        }

        if lineCount == 5 {
            return true
        }
    }

    return false
}

func (b board) score(number int) int {
    sum := 0
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if !b.numbers[x][y].marked {
                sum += b.numbers[x][y].value
            }
        }
    }

    return sum * number
}

func parse_input(filename string) ([]int, []board) {
    f, _ := os.Open(filename)
    scanner := bufio.NewScanner(f)

    scanner.Scan()
    string_numbers := strings.Split(scanner.Text(), ",")
    var numbers []int
    for _, num := range string_numbers {
        num, _ := strconv.Atoi(num)
        numbers = append(numbers, num)
    }

    boards := []board{}
    for scanner.Scan() {
        b := board{}
        for n := 0; n < 5; n++ {
            scanner.Scan()
            string_fields := strings.Fields(scanner.Text())
            for idx, str := range string_fields {
                b.numbers[n][idx].value, _ = strconv.Atoi(str)
            }
        }
        boards = append(boards, b)
    }

    return numbers, boards
}

func play(boards []board, numbers []int) int {
    nBoards := len(boards)
    nWonBoards := 0

    for _, number := range numbers {
        fmt.Println(number)
        for i, _ := range boards {
            if boards[i].mark(number) && !boards[i].hasWon {
                boards[i].hasWon = true
                nWonBoards++

                if nWonBoards == nBoards {
                    return boards[i].score(number)
                }
            }

        }
    }

    return -1
}

func main() {
    numbers, boards := parse_input("input.txt")

    winner := play(boards, numbers)
    fmt.Println(winner)
}
