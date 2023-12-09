package day03

import "testing"

func TestPart1Example(t *testing.T) {
	input := []string{
		"",
	}

	expected := 0

	actual := solve1(input)

	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}

func TestPart1Solution(t *testing.T) {
	expected := 0
	actual := Part1()

	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := []string{
		"",
	}

	expected := 0

	actual := solve2(input)

	if actual != expected {
		t.Errorf("Part2() = %d, expected %d", actual, expected)
	}
}

func TestPart2Solution(t *testing.T) {
	expected := 0
	actual := Part2()

	if actual != expected {
		t.Errorf("Part2() = %d, expected %d", actual, expected)
	}
}
