package main

import "github.com/imperfect-fourth/aoc/day1"

type Day interface {
	Solve()
}

func main() {
	days := []Day{
		day1.New("inputs/day1.txt"),
	}

	for _, d := range days {
		d.Solve()
	}
}
