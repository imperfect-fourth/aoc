package day3

import (
	"errors"
	"fmt"
	"io"
	"os"
	re "regexp"
	"strconv"
)

type day3 struct {
	memory string
}

func New(inputFilepath string) *day3 {
	d := &day3{}
	f, err := os.Open(inputFilepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := d.parseInput(f); err != nil {
		panic(err)
	}
	return d
}

func (d *day3) parseInput(reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	d.memory = string(contents)
	return nil
}

var part1Regex = re.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func (d day3) Part1() string {
	matches := part1Regex.FindAllStringSubmatch(d.memory, -1)
	result := 0
	for _, match := range matches {
		if len(match) != 3 {
			panic(errors.New("unexpected match"))
		}
		a, err := strconv.Atoi(match[1])
		if err != nil {
			panic(fmt.Errorf("unexpected match: %w", err))
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			panic(fmt.Errorf("unexpected match: %w", err))
		}
		result += a * b
	}
	return strconv.Itoa(result)
}

var part2Regex = re.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

func (d day3) Part2() string {
	matches := part2Regex.FindAllStringSubmatch(d.memory, -1)
	result := 0
	do := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if !do {
				continue
			}
			a, err := strconv.Atoi(match[1])
			if err != nil {
				panic(fmt.Errorf("unexpected match: %w", err))
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				panic(fmt.Errorf("unexpected match: %w", err))
			}
			result += a * b
		}
	}
	return strconv.Itoa(result)
}

func (d day3) Day() int {
	return 3
}
