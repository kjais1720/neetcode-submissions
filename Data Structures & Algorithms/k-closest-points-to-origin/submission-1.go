

func kClosest(points [][]int, k int) [][]int {
	QuickSort(points, 0, len(points)-1)
	fmt.Println(points)
	return points[:k]
}

func distanceFromZero(node []int) float64 {
	res := math.Sqrt(math.Pow(float64(node[0]),2) + math.Pow(float64(node[1]),2))
	return res
}

func isNodeSmaller(a, b []int) bool {
	return distanceFromZero(a) < distanceFromZero(b)
}

// [[0,2],[2,0],[2,2]], 0, 2
// left 

func QuickSort(arr [][]int, s, e int) {
	if e-s+1 <= 1 {
		return
	}

	left := s 

	for i := s; i < e ; i++ {
		if isNodeSmaller(arr[i], arr[e]) {
			if i != left{
				temp := arr[i]
				arr[i] = arr[left]
				arr[left] = temp
			}
			left++
		}
	}

	if left == e {
		return
	}

	temp := arr[e]
	arr[e] = arr[left]
	arr[left] = temp 



	QuickSort(arr, s, left-1)
	QuickSort(arr, left+1, e)
}