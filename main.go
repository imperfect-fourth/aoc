package main

import (
	"github.com/imperfect-fourth/aoc/day1"
	"github.com/imperfect-fourth/aoc/day2"
)

type Day interface {
	Solve()
}

func main() {
	days := []Day{
		day1.New("inputs/day1.txt"),
		day2.New("inputs/day2.txt"),
	}

	for _, d := range days {
		d.Solve()
	}
}
