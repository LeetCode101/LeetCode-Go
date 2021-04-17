package algorithms

func numIslands(grid [][]byte) int {
	count := 0
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count += 1

				dfs(grid, i, j, visited)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row int, column int, visited [][]bool) {
	m := len(grid)
	n := len(grid[0])

	if row < 0 || row >= m || column < 0 || column >= n || grid[row][column] != 1 || visited[row][column] {
		return
	}

	visited[row][column] = true

	dfs(grid, row-1, column, visited)
	dfs(grid, row+1, column, visited)
	dfs(grid, row, column-1, visited)
	dfs(grid, row, column+1, visited)
}
