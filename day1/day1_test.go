package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	d := &day1{}
	d.parseInput(strings.NewReader(`1 3
2 3
4 1
7 4
5 4
3 7`))
	assert.Equal(t, d, &day1{left: []int{1, 2, 4, 7, 5, 3}, right: []int{3, 3, 1, 4, 4, 7}})
	assert.Equal(t, d.Part1(), "2")
	assert.Equal(t, d.Part2(), "22")
}
