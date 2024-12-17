package input

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAdjacencyList(t *testing.T) {
	input := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	adjList := GetAdjacencyList(input, [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	})

	for k, v := range adjList {
		switch *k {
		case 0:
			assert.Equal(t, 1, *v[0])
			assert.Equal(t, 3, *v[1])
			assert.Equal(t, 4, *v[2])
		case 1:
			assert.Equal(t, 0, *v[0])
			assert.Equal(t, 2, *v[1])
			assert.Equal(t, 3, *v[2])
			assert.Equal(t, 4, *v[3])
			assert.Equal(t, 5, *v[4])
		case 2:
			assert.Equal(t, 1, *v[0])
			assert.Equal(t, 4, *v[1])
			assert.Equal(t, 5, *v[2])
		case 3:
			assert.Equal(t, 0, *v[0])
			assert.Equal(t, 1, *v[1])
			assert.Equal(t, 4, *v[2])
			assert.Equal(t, 6, *v[3])
			assert.Equal(t, 7, *v[4])
		case 4:
			assert.Equal(t, 0, *v[0])
			assert.Equal(t, 1, *v[1])
			assert.Equal(t, 2, *v[2])
			assert.Equal(t, 3, *v[3])
			assert.Equal(t, 5, *v[4])
			assert.Equal(t, 6, *v[5])
			assert.Equal(t, 7, *v[6])
			assert.Equal(t, 8, *v[7])
		case 5:
			assert.Equal(t, 1, *v[0])
			assert.Equal(t, 2, *v[1])
			assert.Equal(t, 4, *v[2])
			assert.Equal(t, 7, *v[3])
			assert.Equal(t, 8, *v[4])
		case 6:
			assert.Equal(t, 3, *v[0])
			assert.Equal(t, 4, *v[1])
			assert.Equal(t, 7, *v[2])
		case 7:
			assert.Equal(t, 3, *v[0])
			assert.Equal(t, 4, *v[1])
			assert.Equal(t, 5, *v[2])
			assert.Equal(t, 6, *v[3])
			assert.Equal(t, 8, *v[4])
		case 8:
			assert.Equal(t, 4, *v[0])
			assert.Equal(t, 5, *v[1])
			assert.Equal(t, 7, *v[2])
		default:
			assert.FailNow(t, fmt.Sprintf("invalid key: %d", *k))
		}
	}
}
