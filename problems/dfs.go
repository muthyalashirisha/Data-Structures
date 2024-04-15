package problems

import "fmt"

func DistinctIslands(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	mp := map[string]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == 0 && grid[i][j] == 1 {
				path := []byte{}
				dfsIslands(i, j, grid, &path, visited, i, j, n, m)
				fmt.Println(path)
				mp[string(path)]++
			}
		}
	}
	return len(mp)
}
func dfsIslands(row, col int, grid [][]int, path *[]byte, visited [][]int, row0, col0, n, m int) {
	visited[row][col] = 1
	*path = append(*path, byte(row-row0))
	*path = append(*path, byte(col-col0))
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		nr := row + dr[i]
		nc := col + dc[i]
		if nr >= 0 && nr < m && nc >= 0 && nc < n && visited[nr][nc] == 0 && grid[nr][nc] == 1 {
			dfsIslands(nr, nc, grid, path, visited, row0, col0, n, m)
		}
	}
}
