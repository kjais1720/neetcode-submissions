func twoSum(nums []int, target int) []int {
    mp := make(map[int][]int, len(nums))
	for i, n := range nums {
		mp[n] = append(mp[n],i)
	}

	for i, n := range nums {
		if otherIndices, ok := mp[target-n]; ok {
			for _, j := range otherIndices {
				if j != i {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}
