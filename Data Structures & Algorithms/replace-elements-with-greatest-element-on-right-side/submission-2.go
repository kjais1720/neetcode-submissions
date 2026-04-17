func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr)-1; i >=0; i-- {
		if arr[i] < max {
			arr[i] = max
		} else {
			temp := max
			max = arr[i]
			arr[i] = temp
		}
	}
	return arr
}
