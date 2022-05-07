package nodedegree

import (
	"fmt"
)

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if node > nodes {
		return 0, fmt.Errorf("node %v not found in the graph", node)
	}

	var count int
	for _, i := range graph {
		if i[0] == node && i[1] == node {
			count += 2
		}
		if i[0] == node || i[1] == node {
			count++
		}
	}
	return count, nil
}
