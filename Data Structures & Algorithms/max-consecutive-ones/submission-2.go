func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	oneCounter := 0
	for i := 0; i<len(nums); i++ {
		if nums[i] == 1 {
			oneCounter++
		} else {
			oneCounter = 0
		}
		if oneCounter > maxCount {
			maxCount = oneCounter
		}
	}
	return maxCount
}
