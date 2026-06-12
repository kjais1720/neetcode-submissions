func lastStoneWeight(stones []int) int {
	max := 0 
	for _, n := range stones {
		if n > max {
			max = n
		}
	}
	counts := make([]int, max+1)
	for _, n := range stones {
		counts[n]++
	}

	n := max
	for n > 0 {
		counts[n] = counts[n] % 2
		w1 := counts[n] * n
		if w1 == 0 {
			n--
			continue
		}
		counts[n]--
		n--
		for ; n>0; n-- {
			if counts[n] != 0 {
				break
			}
		}
		if n==0{
			return w1
		}

		rem := int(math.Abs(float64(w1-n)))
		counts[rem]++
		counts[n]--
		if rem > n {
			n = rem
		} else if counts[n] <= 0  {
			n--
		} 
	}
	res := 0
	for i := 1; i < len(counts); i++ {
		if counts[i] != 0 {
			res = counts[i] * i
			break
		}
	}
	return res
}
