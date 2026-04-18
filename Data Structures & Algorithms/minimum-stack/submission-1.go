type MinStack struct {
	arr []int
	length int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.arr[this.length-1]
}

func (this *MinStack) GetMin() int {
	min := this.arr[0]
	for i := 1; i< this.length; i++ {
		if this.arr[i] < min {
			min = this.arr[i]
		}
	}
	return min
}
