type DSU struct {
    Parent []int
    Size   []int
}

func (dsu *DSU) find(node int) int {
    if dsu.Parent[node] != node {
        dsu.Parent[node] = dsu.find(dsu.Parent[node])
    }
    return dsu.Parent[node]
}

func (dsu *DSU) union(u, v int) bool {
    pu := dsu.find(u)
    pv := dsu.find(v)
    if pu == pv {
        return false
    }
    if dsu.Size[pu] >= dsu.Size[pv] {
        dsu.Size[pu] += dsu.Size[pv]
        dsu.Parent[pv] = pu
    } else {
        dsu.Size[pv] += dsu.Size[pu]
        dsu.Parent[pu] = pv
    }
    return true
}

func (dsu *DSU) getSize(node int) int {
    par := dsu.find(node)
    return dsu.Size[par]
}

func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    dsu := DSU{
        Parent: make([]int, rows*cols+1),
        Size:   make([]int, rows*cols+1),
    }

    for i := 0; i <= rows*cols; i++ {
        dsu.Parent[i] = i
        dsu.Size[i] = 1
    }

    index := func(r, c int) int {
        return r*cols + c
    }

    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    area := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                for _, dir := range directions {
                    nr, nc := r+dir[0], c+dir[1]
                    if nr < 0 || nc < 0 || nr >= rows ||
                       nc >= cols || grid[nr][nc] == 0 {
                        continue
                    }
                    dsu.union(index(r, c), index(nr, nc))
                }
                area = max(area, dsu.getSize(index(r, c)))
            }
        }
    }
    return area
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}