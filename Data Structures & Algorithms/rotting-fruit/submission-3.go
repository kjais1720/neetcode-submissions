type myq struct {
	arr [][]int
}

func (this *myq) push(n []int) {
	this.arr = append(this.arr, n)
}

func (this *myq) len() int {
	return len(this.arr)
}

func (this *myq) pop() []int {
	if this.len() == 0 {
		return nil
	}
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func orangesRotting(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	mins := 0
	freshCount := 0
	var bfs func()

	rows, cols := len(grid), len(grid[0])
	
	q := &myq{}

	bfs = func() {
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for q.len() > 0 {
			l := q.len()


			for i := 0; i < l; i++ {
				cell := q.pop()
				r, c := cell[0], cell[1]
				if r < 0 || c < 0 || r == rows || c == cols || grid[r][c] == 0 {
					continue
				}


				for _, dir := range dirs {
					newr, newc := r+dir[0], c+dir[1]
					
					if newr < 0 || newc < 0 || newr == rows || newc == cols || grid[newr][newc] == 0 || grid[newr][newc] == 2 {
						continue
					}
					freshCount--
					grid[newr][newc] = 2
					q.push([]int{newr, newc})
					
				}

			}
			if q.len() > 0 {
				mins++
			}
		}
	}

	for i, r := range grid {
		for j, c := range r {
			if c == 1 {
				freshCount++
			} else if c == 2 {
				q.push([]int{i, j})
			}
		}
	}

	bfs()

	if freshCount > 0 {
		return -1
	}
	return mins

}
