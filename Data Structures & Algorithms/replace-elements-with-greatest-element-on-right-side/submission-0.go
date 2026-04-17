func replaceElements(arr []int) []int {
	for i := 0; i <len(arr); i++ {
		if i == len(arr)-1 {
			arr[i] = -1
			break
		}
		max := 0
		for j := i+1; j<len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	return arr
}
