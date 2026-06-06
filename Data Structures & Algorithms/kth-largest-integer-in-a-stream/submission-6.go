
type KthLargest struct {
	arr []int
	kth int
}

func heapify(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	nums = append(nums, nums[0])

	for i := (len(nums) - 1) / 2; i >= 1; i-- {
		idx := i
		left := 2 * idx
		for left < len(nums) {
			if nums[idx] > nums[left] && (left+1 >= len(nums) || nums[left] < nums[left+1]) {
				temp := nums[idx]
				nums[idx] = nums[left]
				nums[left] = temp
				idx = left
				left = idx * 2
			} else if left+1 < len(nums) && nums[idx] > nums[left+1] {
				temp := nums[idx]
				nums[idx] = nums[left+1]
				nums[left+1] = temp
				idx = left + 1
				left = idx * 2
			} else {
				break
			}
		}

	}
	return nums
}

func (this *KthLargest) pop() int {
	if len(this.arr) == 0 {
		return -1
	}
	if len(this.arr) <= 2 {
		res := this.arr[1]
		this.arr = this.arr[:1]
		return res
	}
	res := this.arr[1]
	this.arr[1] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	for i := 1; 2*i < len(this.arr); {
		if this.arr[2*i] < this.arr[i] && (2*i+1 >= len(this.arr) || this.arr[2*i+1] > this.arr[2*i]) {
			temp := this.arr[i]
			this.arr[i] = this.arr[2*i]
			this.arr[2*i] = temp
			i = 2 * i
		} else if 2*i+1 < len(this.arr) && this.arr[2*i+1] < this.arr[i] {
			temp := this.arr[i]
			this.arr[i] = this.arr[2*i+1]
			this.arr[2*i+1] = temp
			i = 2*i + 1
		} else {
			break
		}
	}
	return res
}

func Constructor(k int, nums []int) KthLargest {
	res := KthLargest{
		kth: k,
	}
	sort.Ints(nums)

	if len(nums) < k {
		res.arr = heapify(nums)
	} else {
		res.arr = heapify(nums[len(nums)-k : len(nums)])
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, val, val)
		return val
	}
	if this.arr[1] > val {
		return this.arr[1]
	}
	if len(this.arr) > this.kth {
		this.pop()
	}
	this.arr = append(this.arr, val)
	lastIdx := len(this.arr) - 1
	for i := len(this.arr) / 2; i >= 1 && this.arr[i] > val; i /= 2 {
		temp := this.arr[i]
		this.arr[i] = val
		this.arr[lastIdx] = temp
		lastIdx = i
	}
	return this.arr[1]
}