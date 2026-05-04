func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	s := 0

	e := m * n - 1

	for s<=e {
		mid := (s+e)/2
		midI := mid / n
		midJ := mid % n
		if matrix[midI][midJ] > target {
			e = mid - 1
		} else if matrix[midI][midJ] < target {
			s = mid + 1
		} else {
			return true
		}
	}
	return false
}
