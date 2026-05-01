func sortColors(nums []int) {
    l := 0
	r := len(nums)-1
	i := 0
	for i<=r {
		n := nums[i]
		if n == 0 {
			temp := nums[l]
			nums[l] = 0
			nums[i] = temp
			l++
			i++
		} else if n == 2 {
			temp := nums[r]
			nums[r] = 2
			nums[i] = temp
			r--
		} else {
			i++
		}
	}
}
