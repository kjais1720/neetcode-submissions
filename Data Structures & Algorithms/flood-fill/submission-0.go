func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows := len(image)
	cols := len(image[0])

	if rows == 0 || cols == 0 {
		return image
	}
	originalPixel := image[sr][sc]

	var dfs func(r int, c int)

	dfs = func(r int, c int){
		if r < 0 || c < 0 || r == rows || c == cols || image[r][c] != originalPixel || image[r][c] == color {
			return 
		}

		image[r][c] = color

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)	
	}

	dfs(sr, sc)

	return image
	
}
