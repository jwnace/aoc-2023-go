package day02

import (
	"aoc-2023-go/helpers"
	"math"
	"strings"
)

const RedCubes = 12
const GreenCubes = 13
const BlueCubes = 14

type Game struct {
	id     int
	rounds []Round
}

func (g Game) isPossible() bool {
	for _, round := range g.rounds {
		if round.red > RedCubes || round.green > GreenCubes || round.blue > BlueCubes {
			return false
		}
	}

	return true
}

func (g Game) power() int {
	minRed := 0
	minGreen := 0
	minBlue := 0

	for _, round := range g.rounds {
		minRed = int(math.Max(float64(minRed), float64(round.red)))
		minGreen = int(math.Max(float64(minGreen), float64(round.green)))
		minBlue = int(math.Max(float64(minBlue), float64(round.blue)))
	}

	return minRed * minGreen * minBlue
}

func parseGame(input string) Game {
	parts := strings.Split(input, ": ")
	left := parts[0]
	right := parts[1]

	id := helpers.ToInt(strings.TrimRight(strings.Split(left, " ")[1], ":"))

	rounds := make([]Round, 0)

	for _, round := range strings.Split(right, "; ") {
		rounds = append(rounds, parseRound(round))
	}

	return Game{
		id:     id,
		rounds: rounds,
	}
}
