package day03

import "testing"

func TestPart1Solution(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"part 1 solution", 514969},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2Solution(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"part 2 solution", 78915902},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1Examples(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 example 1", args{[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}}, 4361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve1(tt.args.input); got != tt.want {
				t.Errorf("solve1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 example 1", args{[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}}, 467835},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(tt.args.input); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
