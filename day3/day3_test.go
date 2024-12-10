package day3

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test(t *testing.T) {
	d := &day3{}
	d.parseInput(strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`))
	assert.Equal(t, d, &day3{memory: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`})
	assert.Equal(t, d.Part1(), "161")

	d = &day3{}
	d.parseInput(strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`))
	assert.Equal(t, d.Part2(), "48")
}
