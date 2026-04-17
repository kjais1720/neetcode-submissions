func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	oneCounter := 0
	for i := 0; i<len(nums); i++ {
		if nums[i] == 1 {
			oneCounter++
		}
		if nums[i] != 1 || i == len(nums)-1 {
			if oneCounter > maxCount {
				maxCount = oneCounter
			}
			oneCounter = 0
		}
	}
	return maxCount
}
