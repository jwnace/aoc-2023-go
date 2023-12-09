package day02

import (
	"aoc-2023-go/helpers"
)

var inputLines []string

func init() {
	inputLines = helpers.ReadInputLines()
}

func Part1() int {
	return solve1(inputLines)
}

func Part2() int {
	return solve2(inputLines)
}

func solve1(input []string) int {
	games := make([]Game, 0)

	for _, line := range input {
		games = append(games, parseGame(line))
	}

	sum := 0

	for _, game := range games {
		if game.isPossible() {
			sum += game.id
		}
	}

	return sum
}

func solve2(input []string) int {
	games := make([]Game, 0)

	for _, line := range input {
		games = append(games, parseGame(line))
	}

	sum := 0

	for _, game := range games {
		sum += game.power()
	}

	return sum
}
