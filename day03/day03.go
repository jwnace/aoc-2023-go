package day03

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
	grid := buildGrid(input)
	partNumbers := getPartNumbers(grid)

	return sumPartNumbers(partNumbers)
}

func solve2(input []string) int {
	grid := buildGrid(input)
	partNumbers := getPartNumbers(grid)

	return sumGearRatios(grid, partNumbers)
}

func buildGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))

	for row, line := range input {
		grid[row] = make([]rune, len(line))
		for col, char := range line {
			grid[row][col] = char
		}
	}

	return grid
}

func getPartNumbers(grid [][]rune) []PartNumber {
	partNumbers := make([]PartNumber, 0)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			currentNumber, start, end := 0, 0, 0

			for ; col < len(grid[row]); col++ {
				if grid[row][col] < '0' || grid[row][col] > '9' {
					break
				}

				if currentNumber == 0 {
					start = col
					end = col
				} else {
					end = col
				}

				currentNumber = currentNumber*10 + int(grid[row][col]-'0')
			}

			if currentNumber > 0 {
				for r := row - 1; r <= row+1; r++ {
					for c := start - 1; c <= end+1; c++ {
						if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
							continue
						}

						if grid[r][c] != '.' && (grid[r][c] < '0' || grid[r][c] > '9') {
							partNumbers = append(partNumbers, PartNumber{currentNumber, row, start, end})
						}
					}
				}
			}
		}
	}

	return partNumbers
}

func sumPartNumbers(numbers []PartNumber) int {
	sum := 0

	for _, number := range numbers {
		sum += number.number
	}

	return sum
}

func sumGearRatios(grid [][]rune, partNumbers []PartNumber) int {
	result := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '*' {
				count := 0
				adjacentNumbers := make([]int, 0)

				for _, number := range partNumbers {
					if number.row >= row-1 && number.row <= row+1 && number.endCol >= col-1 && number.startCol <= col+1 {
						count++
						adjacentNumbers = append(adjacentNumbers, number.number)
					}
				}

				if count == 2 {
					result += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}

	return result
}
