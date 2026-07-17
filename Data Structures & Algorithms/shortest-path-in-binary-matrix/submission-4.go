func shortestPathBinaryMatrix(grid [][]int) int {

	rows, cols := len(grid), len(grid[0])
	cells := make([][]int, 0, rows*cols)

	length := 0

	cells = append(cells, []int{0, 0})

	r, c := cells[0][0], cells[0][1]
	if grid[r][c] == 0 {
		length++
	} else {
		return -1
	}
	if grid[rows-1][cols-1] != 0 {
		return -1
	}

	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	l := len(cells)
	counted := false
	for i := 0; i < l; i++ {

		r, c = cells[i][0], cells[i][1]
		if r == rows-1 && c == cols-1 {
			return length
		}
		for _, dir := range dirs {
			dr, dc := dir[0], dir[1]
			newr, newc := r+dr, c+dc
			if newr < 0 || newc < 0 || newr == rows || newc == cols || grid[newr][newc] == 1 {
				if newr == rows-1 && newc == cols-1 && grid[newr][newc] == 1 {
					return -1
				}
				continue
			}
			

			cells = append(cells, []int{newr, newc})
			
			grid[newr][newc] = 1

			if !counted {
				length++
				counted = true
			}
			if newr == rows-1 && newc == cols-1 {
				return length
			}
		}
		if i == l-1 {
			l = len(cells)
			counted = false
		}

	}
	if length == 0 || r != rows-1 || c != cols-1 {
		return -1
	}
	return length

}
