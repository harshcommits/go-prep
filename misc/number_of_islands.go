package misc

func NumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && !visited[r][c] {
				count++
				exploreIsland(grid, visited, r, c, rows, cols)
			}
		}
	}

	return count
}

func exploreIsland(grid [][]byte, visited [][]bool, row, col, rows, cols int) {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return
	}
	if grid[row][col] != '1' || visited[row][col] {
		return
	}

	visited[row][col] = true

	exploreIsland(grid, visited, row-1, col, rows, cols)
	exploreIsland(grid, visited, row+1, col, rows, cols)
	exploreIsland(grid, visited, row, col-1, rows, cols)
	exploreIsland(grid, visited, row, col+1, rows, cols)
}
