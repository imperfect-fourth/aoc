package day2

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test(t *testing.T) {
	d := &day2{}
	d.parseInput(strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`))
	assert.Equal(t, d, &day2{
		reports: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	})
	assert.Equal(t, d.Part1(), "2")
	assert.Equal(t, d.Part2(), "4")
}
