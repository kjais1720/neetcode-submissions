var res = map[int]int{
	1:1,
	2:2,
}

func climbStairs(n int) int {
	if n<=0 {
		return 0
	}
	if val, ok := res[n]; ok {
		return val
	}

	val := climbStairs(n-1) + climbStairs(n-2)
	res[n] = val
	return val
}
