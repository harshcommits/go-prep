package misc

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func NumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && !visited[r][c] {
				count++
				dfs(grid, visited, r, c, rows, cols)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, visited [][]bool, row, col, rows, cols int) {
	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != '1' || visited[row][col] {
		return
	}

	visited[row][col] = true
	for _, d := range directions {
		dfs(grid, visited, row+d[0], col+d[1], rows, cols)
	}
}
