package day01

import (
    "aoc-2023-go/helpers"
    "strings"
)

var inputFile string
var digitStrings map[string]int

func init() {
    inputFile = helpers.ReadInput()

    digitStrings = make(map[string]int)
    digitStrings["one"] = 1
    digitStrings["two"] = 2
    digitStrings["three"] = 3
    digitStrings["four"] = 4
    digitStrings["five"] = 5
    digitStrings["six"] = 6
    digitStrings["seven"] = 7
    digitStrings["eight"] = 8
    digitStrings["nine"] = 9
}

func Part1() int {
    return solve1(inputFile)
}

func Part2() int {
    return solve2(inputFile)
}

func solve1(input string) int {
    sum := 0

    for _, line := range strings.Split(input, "\n") {
        digits := make([]int, 0)

        for _, c := range line {
            if c >= '0' && c <= '9' {
                digits = append(digits, int(c)-'0')
            }
        }

        sum += digits[0]*10 + digits[len(digits)-1]
    }

    return sum
}

func solve2(input string) int {
    sum := 0

    for _, line := range strings.Split(input, "\n") {
        digits := make([]int, 0)

        for i, c := range line {
            if c >= '0' && c <= '9' {
                digits = append(digits, int(c)-'0')
            }

            for key, value := range digitStrings {
                if strings.HasPrefix(line[i:], key) {
                    digits = append(digits, value)
                }
            }
        }

        sum += digits[0]*10 + digits[len(digits)-1]
    }

    return sum
}
