func getConcatenation(nums []int) []int {
	n := len(nums)
    ans := make([]int, 2*n)
	for i := 0; i< 2*n; i++ {
		idx := i % n
		ans[i] = nums[idx]
	}
	return ans
}
