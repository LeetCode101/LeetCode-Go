package algorithms

/**
Must change grid[i][j] == 1 to grid[i][j] == '1' on LeetCode.
*/
func numIslands1(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

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
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	if row < 0 || row >= m || column < 0 || column >= n || grid[row][column] != 1 || visited[row][column] {
		return
	}

	visited[row][column] = true

	for _, direction := range directions {
		dfs(grid, row+direction[0], column+direction[1], visited)
	}
}
