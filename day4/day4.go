package day4

import (
	"bytes"
	"io"
	"os"
	"strconv"
)

type day4 struct {
	graph [][]byte
}

func (d day4) Day() int {
	return 4
}

func New(inputFilepath string) *day4 {
	d := &day4{}
	f, err := os.Open(inputFilepath)
	if err != nil {
		panic(err)
	}
	if err := d.parseInput(f); err != nil {
		panic(err)
	}
	return d
}

func (d *day4) parseInput(reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	lines := bytes.Split(contents, []byte("\n"))
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	d.graph = lines
	return nil
}

func (d day4) Part1() string {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	count := 0
	for i := range len(d.graph) {
		for j := range len(d.graph[0]) {
			if rune(d.graph[i][j]) != 'X' {
				continue
			}
			for _, dir := range directions {
				count += checkDirection(d.graph, dir, i, j)
			}
		}
	}
	return strconv.Itoa(count)
}

func checkDirection(graph [][]byte, dir [2]int, i, j int) int {
	if i+(3*dir[0]) < 0 || i+(3*dir[0]) >= len(graph) {
		return 0
	}
	if j+(3*dir[1]) < 0 || j+(3*dir[1]) >= len(graph[0]) {
		return 0
	}

	currLetter := 'X'
	for {
		next := nextLetter(currLetter)
		i += dir[0]
		j += dir[1]
		if rune(graph[i][j]) != next {
			return 0
		}
		if next == 'S' {
			return 1
		}
		currLetter = next
	}
	return 0
}

//		adjList := input.GetAdjacencyList(d.graph, [][2]int{
//			{-1, -1}, {-1, 0}, {-1, 1},
//			{0, -1}, {0, 1},
//			{1, -1}, {1, 0}, {1, 1},
//		})
//
//		var keys []*byte
//		c := 0
//		for k := range adjList {
//			if *k == byte('X') {
//				c++
//			}
//			keys = append(keys, k)
//		}
//		fmt.Println(c)
//		xmasCount := countXmasFromLetter(keys, 'X', adjList, "")
//		return strconv.Itoa(xmasCount)
//	}
func (d day4) Part2() string {
	count := 0
	for i := range len(d.graph) {
		if i == 0 || i == len(d.graph)-1 {
			continue
		}
		for j := range len(d.graph[0]) {
			if j == 0 || j == len(d.graph[0])-1 {
				continue
			}
			if rune(d.graph[i][j]) != 'A' {
				continue
			}
			if tl := rune(d.graph[i-1][j-1]); tl != 'M' && tl != 'S' {
				continue
			} else if br := rune(d.graph[i+1][j+1]); br != 'M' && br != 'S' {
				continue
			} else if tl == br {
				continue
			} else if tr := rune(d.graph[i-1][j+1]); tr != 'M' && tr != 'S' {
				continue
			} else if bl := rune(d.graph[i+1][j-1]); bl != 'M' && bl != 'S' {
				continue
			} else if tr == bl {
				continue
			}
			count++
		}
	}
	return strconv.Itoa(count)
}

//	func countXmasFromLetter(searchList []*byte, letter rune, adjList map[*byte][]*byte, depth string) int {
//		fmt.Printf("%ssearching %s\n", depth, string(letter))
//		count := 0
//		for _, x := range searchList {
//			fmt.Printf("%s  search %s\n", depth, string([]byte{*x}))
//			if *x == byte(letter) {
//				if letter == 'S' {
//					return 1
//				}
//				count += countXmasFromLetter(adjList[x], nextLetter(letter), adjList, fmt.Sprintf("%s  ", depth))
//			}
//		}
//		fmt.Printf("%s%s %d\n", depth, string(letter), count)
//		return count
//	}
func nextLetter(letter rune) rune {
	switch letter {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	}
	return 'S'
}
