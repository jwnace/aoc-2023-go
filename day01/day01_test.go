package day01

import "testing"

func TestPart1Example(t *testing.T) {
    expected := 142
    actual := solve1("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")

    if actual != expected {
        t.Errorf("Part1() = %d, expected %d", actual, expected)
    }
}

func TestPart1Solution(t *testing.T) {
    expected := 55208
    actual := Part1()

    if actual != expected {
        t.Errorf("Part1() = %d, expected %d", actual, expected)
    }
}

func TestPart2Example(t *testing.T) {
    expected := 281
    actual := solve2("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")

    if actual != expected {
        t.Errorf("Part2() = %d, expected %d", actual, expected)
    }
}

func TestPart2Solution(t *testing.T) {
    expected := 54578
    actual := Part2()

    if actual != expected {
        t.Errorf("Part2() = %d, expected %d", actual, expected)
    }
}
