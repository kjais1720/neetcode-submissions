func search(nums []int, target int) int {
	s := 0
	e := len(nums)-1
	for s <= e {
		mid := (s+e)/2
		if target < nums[mid] {
			e = mid-1
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
