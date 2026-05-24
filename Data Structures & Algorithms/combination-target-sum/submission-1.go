func combinationSum(nums []int, target int) [][]int {
    subs := make([][]int, 0)
	sub := make([]int, 0, len(nums))

	var dfs func([]int, int) 
	dfs = func(nm []int, sm int){
		if sm >= target || len(nm) == 0  {
			if sm > target || len(nm) == 0 {
				return
			} else {
				subCopy := make([]int, len(sub))
				copy(subCopy, sub)
				subs = append(subs, subCopy)
				return
			}
		}
		sub = append(sub, nm[0])
		sm += nm[0]
		dfs(nm, sm)
		sm -= nm[0]
		sub = sub[:len(sub)-1]
		dfs(nm[1:], sm)
	}
	dfs(nums, 0)
	return subs
}

func arrSum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}