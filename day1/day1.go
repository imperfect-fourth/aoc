package day1

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

type day1 struct {
	left  []int
	right []int
}

func New(inputFilepath string) *day1 {
	d := &day1{}
	f, err := os.Open(inputFilepath)
	if err != nil {
		panic(err)
	}
	if err := d.parseInput(f); err != nil {
		panic(err)
	}
	return d
}

func (d *day1) parseInput(reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	lines := bytes.Split(contents, []byte("\n"))
	d.left = make([]int, len(lines))
	d.right = make([]int, len(lines))
	for i, line := range lines {
		fields := bytes.Fields(line)
		if i == len(lines)-1 && len(fields) == 0 {
			return nil
		}
		if len(fields) != 2 {
			return errors.New("invalid input: expected exactly 2 fields in a line")
		}

		l, err := strconv.Atoi(string(fields[0]))
		if err != nil {
			return fmt.Errorf("invalid input: %w", err)
		}
		d.left[i] = l
		r, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			return fmt.Errorf("invalid input: %w", err)
		}
		d.right[i] = r
	}
	return nil
}

func (d day1) Part1() string {
	sortedLeft := slices.Sorted(slices.Values(d.left))
	sortedRight := slices.Sorted(slices.Values(d.right))

	distance := 0
	for i := range len(sortedLeft) {
		s := sortedLeft[i] - sortedRight[i]
		if s < 0 {
			distance -= s
			continue
		}
		distance += s
	}
	return strconv.Itoa(distance)
}

func (d day1) Part2() string {
	frequency := make(map[int]int)
	for _, n := range d.right {
		frequency[n] += 1
	}

	similarity := 0
	for _, n := range d.left {
		similarity += n * frequency[n]
	}
	return strconv.Itoa(similarity)
}

func (d day1) Day() int {
	return 1
}
