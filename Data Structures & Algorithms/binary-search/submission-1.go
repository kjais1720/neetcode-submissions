func binary(nums[]int, target, s, e int) int {
	if s > e {
		return -1
	}

	mid := (s+e)/2
	if nums[mid] > target {
		return binary(nums, target, s, mid-1)
	} else if nums[mid] < target {
		return binary(nums, target, mid+1, e)
	} else {
		return mid
	}
}

func search(nums []int, target int) int {
	return binary(nums, target, 0, len(nums)-1)
}
