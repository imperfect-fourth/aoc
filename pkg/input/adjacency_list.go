package input

func GetAdjacencyList[T any](graph [][]T, directions [][2]int) map[*T][]*T {
	rows := len(graph)
	cols := len(graph[0])
	nodes := make(map[[2]int]*T)
	adjacencyList := make(map[*T][]*T)

	for r := range rows {
		for c := range cols {
			indices := [2]int{r, c}
			val := graph[r][c]
			if nodes[indices] == nil {
				nodes[indices] = &val
			}
			node := nodes[indices]
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					neighbourIndices := [2]int{nr, nc}
					val := graph[nr][nc]
					if nodes[neighbourIndices] == nil {
						nodes[neighbourIndices] = &val
					}
					neighbor := nodes[neighbourIndices]
					adjacencyList[node] = append(adjacencyList[node], neighbor)
				}
			}
		}
	}
	return adjacencyList
}
