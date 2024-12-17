package day5

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	val   string
	level *int
}

type day5 struct {
	nodeMap map[string]*node
	adjList map[*node][]*node

	ops [][]string
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
	makeGraph(d, split[0])
	d.ops = splitOps(split[1])

	for val, node := range d.nodeMap {
		fmt.Println(val, *(node.level))
	}
	return nil
}

func (d day5) Part1() string {
	ans := 0
	for _, op := range d.ops {
		currLevel := math.MinInt
		ok := true
		for _, update := range op {
			if _, nok := d.nodeMap[update]; !nok {
				ok = false
				fmt.Println("not found")
				break
			}
			if *(d.nodeMap[update].level) < currLevel {
				fmt.Println(op)
				fmt.Println(*(d.nodeMap[update]))
				ok = false
				break
			}
			currLevel = *(d.nodeMap[update].level)
		}
		if ok {
			fmt.Println(op)
			mid, _ := strconv.Atoi(string(op[(len(op)-1)/2]))
			ans += mid
		}
	}
	return strconv.Itoa(ans)
}

func splitOps(raw []byte) [][]string {
	rows := bytes.Split(raw, []byte("\n"))
	ops := make([][]string, len(rows))
	for i, row := range rows {
		ops[i] = strings.Split(string(row), ",")
	}
	return ops
}

func makeGraph(d *day5, graph []byte) {
	edges := bytes.Split(graph, []byte("\n"))
	nodeMap := make(map[string]*node)
	adjList := make(map[*node][]*node)

	delimeter := []byte("|")
	for _, edge := range edges {
		nodes := bytes.Split(edge, delimeter)
		node0 := string(nodes[0])
		node1 := string(nodes[1])
		if _, ok := nodeMap[node1]; !ok {
			nodeMap[node1] = &node{val: node1}
		}
		if _, ok := nodeMap[node0]; !ok {
			nodeMap[node0] = &node{val: node0}
		}
		adjList[nodeMap[node0]] = append(adjList[nodeMap[node0]], nodeMap[node1])
	}

	for k, _ := range adjList {
		assignLevel(k, adjList, 0, "")
	}
	d.nodeMap = nodeMap
	d.adjList = adjList
}

func assignLevel(n *node, adjList map[*node][]*node, level int, tab string) int {
	fmt.Printf("%sassigning %s %d\n", tab, n.val, level)
	if n.level != nil {
		return *(n.level)
	}
	minLevel := level
	for _, neighbour := range adjList[n] {
		if neighbour.level == nil {
			continue
		}
		fmt.Printf("%schild set %s %d\n", tab, neighbour.val, *(neighbour.level))
		minLevel = min(minLevel, *(neighbour.level)-1)
	}
	n.level = &minLevel
	fmt.Printf("%sassigned %s %d\n", tab, string(n.val), *(n.level))
	for _, neighbour := range adjList[n] {
		if neighbour.level != nil {
			continue
		}
		assigned := assignLevel(neighbour, adjList, *(n.level)+1, fmt.Sprintf("%s  ", tab))
		if assigned <= *(n.level) {
			*(n.level) = assigned - 1
			fmt.Printf("%schanged %s %d\n", tab, string(n.val), *(n.level))
		}
	}
	return *(n.level)
}

func (d day5) Part2() string {
	return ""
}

func (d day5) Day() int {
	return 5
}
