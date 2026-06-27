func heappify(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} 
	nums = append(nums, nums[0])
	if len(nums) < 2 {
		return nums
	}

	for i := len(nums) / 2; i > 0 ; i-- {
		j := i
		for 2 * j < len(nums) {
			if nums[2*j] < nums[j] && ((2 * j + 1) >= len(nums) || nums[2 * j + 1] > nums[2*j]) {
				temp := nums[j]
				nums[j] = nums[2*j]
				nums[2*j] = temp
				j *= 2
			} else if (2 * j + 1) < len(nums) && nums[2 * j + 1] < nums[j] {
				temp := nums[j]
				nums[j] = nums[2*j+1]
				nums[2*j+1] = temp
				j = 2*j+1 
			} else {
				break
			}
		}
	}

	return nums
}

type heapp struct {
	nums []int
}

func (h *heapp) len() int {
	return len(h.nums)
}

func (h *heapp) push(n int) {
	if h.len() < 1 {
		h.nums = append(h.nums, n, n)
		return
	}
	h.nums = append(h.nums, n)
	l := h.len() /2
	m := h.len()-1
	for l >= 0 && h.nums[l] > h.nums[m] {
		temp := h.nums[m]
		h.nums[m] = h.nums[l]
		h.nums[l] = temp
		m = l
		l /= 2
	}
}

func (h *heapp) pop() int {
	if h.len() == 0 {
		return -1
	} else if h.len() == 2{
		res := h.nums[1]
		h.nums = h.nums[:1]
		return res
	}
	res := h.nums[1]
	h.nums[1] = h.nums[h.len()-1]
	h.nums = h.nums[:h.len()-1]

	for i := 1; 2*i < h.len(); {
		if h.nums[2*i] < h.nums[i] && ((2*i+1) >= h.len() || h.nums[2*i+1] > h.nums[2*i]) {
			temp := h.nums[i]
			h.nums[i] = h.nums[2*i]
			h.nums[2*i] = temp
			i = 2*i
		} else if (2*i+1) < h.len() && h.nums[2*i+1] < h.nums[i] {
			temp := h.nums[i]
			h.nums[i] = h.nums[2*i+1]
			h.nums[2*i+1] = temp
			i = 2*i+1
		} else {
			break
		}
	}
	return res
}

func NewHeapp(nums []int) heapp {
	nums = heappify(nums)
	return heapp{
		nums: nums,
	}
}

func findKthLargest(nums []int, k int) int {
	h := NewHeapp(nums)
	for i := 0 ; i < len(nums) - k; i++ {
		h.pop()
	}
	return h.pop()
}
