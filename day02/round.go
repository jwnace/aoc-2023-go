package day02

import (
    "aoc-2023-go/helpers"
    "strings"
)

type Round struct {
    red   int
    green int
    blue  int
}

func parseRound(round string) Round {
    red := 0
    green := 0
    blue := 0

    for _, cube := range strings.Split(round, ", ") {
        if strings.HasSuffix(cube, "red") {
            red = helpers.ToInt(strings.TrimSuffix(cube, " red"))
        } else if strings.HasSuffix(cube, "green") {
            green = helpers.ToInt(strings.TrimSuffix(cube, " green"))
        } else if strings.HasSuffix(cube, "blue") {
            blue = helpers.ToInt(strings.TrimSuffix(cube, " blue"))
        }
    }

    return Round{
        red:   red,
        green: green,
        blue:  blue,
    }
}
