package main

import (
	"fmt"

	"github.com/imperfect-fourth/aoc/day1"
	"github.com/imperfect-fourth/aoc/day2"
	"github.com/imperfect-fourth/aoc/day3"
	"github.com/imperfect-fourth/aoc/day4"
)

type Day interface {
	Day() int
	Part1() string
	Part2() string
}

func main() {
	days := []Day{
		day1.New("inputs/day1.txt"),
		day2.New("inputs/day2.txt"),
		day3.New("inputs/day3.txt"),
		day4.New("inputs/day4.txt"),
	}

	for _, d := range days {
		Solve(d)
	}
}

func Solve(d Day) {
	fmt.Printf("Day %d Part 1: %s\n", d.Day(), d.Part1())
	fmt.Printf("Day %d Part 2: %s\n", d.Day(), d.Part2())
}
