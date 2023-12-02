package day02

import "testing"

func TestPart1Example(t *testing.T) {
    input := []string{
        "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
        "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
        "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    }

    expected := 8

    actual := solve1(input)

    if actual != expected {
        t.Errorf("Part1() = %d, expected %d", actual, expected)
    }
}

func TestPart1Solution(t *testing.T) {
    expected := 2348
    actual := Part1()

    if actual != expected {
        t.Errorf("Part1() = %d, expected %d", actual, expected)
    }
}

func TestPart2Example(t *testing.T) {
    input := []string{
        "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
        "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
        "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    }

    expected := 2286

    actual := solve2(input)

    if actual != expected {
        t.Errorf("Part2() = %d, expected %d", actual, expected)
    }
}

func TestPart2Solution(t *testing.T) {
    expected := 76008
    actual := Part2()

    if actual != expected {
        t.Errorf("Part2() = %d, expected %d", actual, expected)
    }
}
