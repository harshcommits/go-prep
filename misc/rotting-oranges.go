package misc

func OrangesRotting(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	queue := make([][2]int, 0)
	freshCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	for len(queue) > 0 && freshCount > 0 {
		minutes++
		currentLevelSize := len(queue)

		for i := 0; i < currentLevelSize; i++ {
			orange := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				newRow := orange[0] + dir[0]
				newCol := orange[1] + dir[1]

				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					freshCount--
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
	}

	if freshCount > 0 {
		return -1
	}

	return minutes
}
