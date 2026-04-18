func calPoints(operations []string) int {
	scores := make([]int, 0, len(operations))
	for _, i  := range operations {
		if score, err := strconv.Atoi(i); err == nil {
			scores = append(scores, score)
		} else if i == "+" {
			scores = append(scores, scores[len(scores)-1] + scores[len(scores)-2])
		} else if i == "C" {
			scores = scores[0:len(scores)-1]
		} else if i == "D" {
			scores = append(scores, 2*scores[len(scores)-1])
		}
	}
	sum := 0
	for _, j := range scores {
		sum += j
	}
	return sum
}
