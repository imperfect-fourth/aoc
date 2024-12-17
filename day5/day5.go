package day5

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	val   string
	level *int
}

type day5 struct {
	ordering map[string]map[string]int
	ops      [][]string
}

func New(inputFilepath string) *day5 {
	d := &day5{}
	f, err := os.Open(inputFilepath)
	if err != nil {
		panic(err)
	}
	if err := d.parseInput(f); err != nil {
		panic(err)
	}
	return d
}

func (d *day5) parseInput(reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	split := bytes.Split(contents, []byte("\n\n"))
	if len(split) != 2 {
		return errors.New("invalid input")
	}
	d.ordering = getOrdering(split[0])
	d.ops = splitOps(split[1])
	return nil
}

func (d day5) Part1() string {
	ans := 0
	for _, op := range d.ops {
		ok := true
	opLoop:
		for i := 0; i < len(op)-1; i++ {
			if _, nok := d.ordering[op[i]]; !nok {
				ok = false
				break
			}
			for j := i + 1; j < len(op); j++ {
				if d.ordering[op[i]][op[j]] != 1 {
					ok = false
					break opLoop
				}
			}
		}
		if ok {
			mid, _ := strconv.Atoi(string(op[(len(op)-1)/2]))
			ans += mid
		}
	}
	return strconv.Itoa(ans)
}

func (d day5) Part2() string {
	ans := 0
	for _, op := range d.ops {
		ok := true
	opLoop:
		for i := 0; i < len(op)-1; i++ {
			if _, nok := d.ordering[op[i]]; !nok {
				ok = false
				break
			}
			for j := i + 1; j < len(op); j++ {
				if d.ordering[op[i]][op[j]] != 1 {
					ok = false
					break opLoop
				}
			}
		}
		if !ok {
			fmt.Println(op)
			slices.SortFunc(op, func(a, b string) int {
				if _, ok := d.ordering[a]; ok {
					if _, ok := d.ordering[a][b]; ok {
						return 1
					}
					return -1
				}
				return -1
			})
			fmt.Println(op)
			mid, _ := strconv.Atoi(string(op[(len(op)-1)/2]))
			ans += mid
		}
	}
	return strconv.Itoa(ans)
}

func (d day5) Day() int {
	return 5
}

func splitOps(raw []byte) [][]string {
	rows := bytes.Split(raw, []byte("\n"))
	ops := make([][]string, len(rows))
	for i, row := range rows {
		ops[i] = strings.Split(string(row), ",")
	}
	return ops
}

func getOrdering(graph []byte) map[string]map[string]int {
	orderings := bytes.Split(graph, []byte("\n"))
	ordering := make(map[string]map[string]int)

	delimeter := []byte("|")
	for _, o := range orderings {
		n := bytes.Split(o, delimeter)
		n0 := string(n[0])
		n1 := string(n[1])

		if _, ok := ordering[n0]; !ok {
			ordering[n0] = make(map[string]int)
		}
		ordering[n0][n1] = 1
	}
	return ordering
}
