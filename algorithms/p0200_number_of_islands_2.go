package algorithms

/**
Must change grid[i][j] == 1 to grid[i][j] == '1' on LeetCode.
*/
func numIslands2(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	count := 0
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	queue := make([][2]int, 0, n)
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count += 1
				visited[i][j] = true
				queue = append(queue, [2]int{i, j})

				for len(queue) > 0 {
					row, column := queue[0][0], queue[0][1]
					queue = queue[1:]

					for _, direction := range directions {
						nextRow, nextColumn := row+direction[0], column+direction[1]

						if !isValid(grid, nextRow, nextColumn, visited) {
							continue
						}

						queue = append(queue, [2]int{nextRow, nextColumn})
						visited[nextRow][nextColumn] = true
					}
				}
			}
		}
	}

	return count
}

func isValid(grid [][]byte, row int, column int, visited [][]bool) bool {
	m := len(grid)
	n := len(grid[0])

	return !(row < 0 || row >= m || column < 0 || column >= n || grid[row][column] != 1 || visited[row][column])
}
