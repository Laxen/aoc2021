package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type board struct {
    numbers [5][5] int
}

func (b *board) show() {
    var sb strings.Builder
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            sb.WriteString(strconv.Itoa(b.numbers[x][y]))
            sb.WriteString("\t")
        }
        sb.WriteString("\n")
    }

    fmt.Println(sb.String())
}

func main() {
    f, _ := os.Open("example_input.txt")
    scanner := bufio.NewScanner(f)

    scanner.Scan()
    string_numbers := strings.Split(scanner.Text(), ",")
    var numbers []int
    for _, num := range string_numbers {
        num, _ := strconv.Atoi(num)
        numbers = append(numbers, num)
    }

    boards := make([]board, 3)
    for i := 0; i < 3; i++ {
        scanner.Scan()
        for n := 0; n < 5; n++ {
            scanner.Scan()
            string_fields := strings.Fields(scanner.Text())
            for idx, str := range string_fields {
                boards[i].numbers[n][idx], _ = strconv.Atoi(str)
            }
        }
    }
}
