// import "fmt"
func removeElement(nums []int, val int) int {
    remaining := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			cnt++
		} else {
			nums[i-cnt] = nums[i]
			remaining++
		}
	}
	fmt.Println(nums)
	return remaining
}

