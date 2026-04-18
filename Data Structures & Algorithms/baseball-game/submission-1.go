func calPoints(operations []string) int {
	scores := make([]int, 0, len(operations))
	sum := 0
	for _, i  := range operations {
		if score, err := strconv.Atoi(i); err == nil {
			sum += score
			scores = append(scores, score)
		} else if i == "+" {
			scores = append(scores, scores[len(scores)-1] + scores[len(scores)-2])
			sum += scores[len(scores)-1]
		} else if i == "C" {
			sum -= scores[len(scores)-1]
			scores = scores[0:len(scores)-1]
		} else if i == "D" {
			sum += 2 * scores[len(scores)-1]
			scores = append(scores, 2*scores[len(scores)-1])
		}
	}
	return sum
}
