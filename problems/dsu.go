package problems

import "fmt"

type DSU struct {
	parent, rank []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &DSU{parent, rank}
}

func (dsu *DSU) find(u int) int {
	if dsu.parent[u] != u {
		dsu.parent[u] = dsu.find(dsu.parent[u])
	}
	return dsu.parent[u]
}

func (dsu *DSU) union(u, v int) {
	rootU := dsu.find(u)
	rootV := dsu.find(v)

	if rootU != rootV {
		if dsu.rank[rootU] > dsu.rank[rootV] {
			dsu.parent[rootV] = rootU
		} else if dsu.rank[rootU] < dsu.rank[rootV] {
			dsu.parent[rootU] = rootV
		} else {
			dsu.parent[rootV] = rootU
			dsu.rank[rootU]++
		}
	}
}

func HasCycle(graph [][]int, V int) bool {
	dsu := newDSU(V)
	fmt.Println(dsu)

	for _, edge := range graph {
		u, v := edge[0], edge[1]
		fmt.Println(u, v)
		rootU, rootV := dsu.find(u), dsu.find(v)
		fmt.Println(rootU, rootV)
		if rootU == rootV {
			return true // Cycle detected
		}

		dsu.union(rootU, rootV)
	}

	return false // No cycle detected
}
