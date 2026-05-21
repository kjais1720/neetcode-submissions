func subsets(nums []int) [][]int {
	var dfs func( i int)

	subs := make([][]int, 0)
	subset := make([]int, 0)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			subs = append(subs, temp)
			return
		}
		subset = append(subset, nums[i])
		dfs(i+1)

		subset = subset[0: len(subset)-1]
		dfs(i+1)

	}
	dfs(0)
	return subs
}
