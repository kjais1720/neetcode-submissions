func maxAreaOfIsland(grid [][]int) int {
    var dfs func(grid [][]int, r, c int)
	rows, cols := len(grid), len(grid[0])
	max, count := 0, 0
	dfs = func(grid [][]int, r, c int) {
		if r < 0 || c < 0 || r == rows || c == cols || grid[r][c] == 0 {
			return
		}
		count++
		grid[r][c] = 0
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				dfs(grid, i, j)
				if count > max {
					max = count
				}
				count = 0
			}
		}
	}

	return max
}
