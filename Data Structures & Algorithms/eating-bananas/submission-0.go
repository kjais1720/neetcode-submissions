func minEatingSpeed(piles []int, h int) int {
	rate := 1
	lastRate := 1
	
	for {
		if isRateEnough(piles, rate, h) {
			temp := (rate + lastRate) / 2
			fmt.Println(rate, lastRate)
			if temp <= lastRate {
				break
			} else {
				rate = temp
			}
		} else {
			fmt.Println("rate",rate)
			lastRate = rate
			rate *= 2
		}
	}

	return rate
}

func isRateEnough (piles []int, rate, hours int) bool {
	res := 0 
	for _, b := range piles {
		res += b / rate
		if b % rate != 0 {
			res += 1
		}
		if res > hours {
			return false
		}
	}

	return true
}