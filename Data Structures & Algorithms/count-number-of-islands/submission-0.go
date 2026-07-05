func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	visits := make(map[int]map[int]bool, rows)

	var dfs func(grid [][]byte, r, c int) bool

	dfs = func(grid [][]byte, r, c int) bool {
		rvisits, ok := visits[r]

		if r < 0 || c < 0 || r == rows || c == cols || grid[r][c] == '0' || (ok && rvisits[c] == true) {
			return true
		}

		if ok {
			rvisits[c] = true
		} else {
			visits[r] = map[int]bool{
				c: true,
			}
		}

		if dfs(grid, r+1, c) && dfs(grid, r-1, c) && dfs(grid, r, c+1) && dfs(grid, r, c-1) {
			return true
		}
		return false
	}

	count := 0
	for i, row := range grid {
		for j, val := range row {
			if rvisits, ok := visits[i]; val == '0' || (ok && rvisits[j] == true) {
				continue
			}
			if dfs(grid, i, j) {
				count++
			}
		}
	}

	return count
}