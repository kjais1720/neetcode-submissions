func sortColors(nums []int) {
    counts := []int{0,0,0}
	for _, n := range nums {
		counts[n]++
	}
	fmt.Println(counts)
	i := 0
	for idx, count := range counts {
		for range count {
			nums[i] = idx
			i++
		}
	}
}
