type MinStack struct {
	arr []int
	length int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.length == 0 {
		this.mins = append(this.mins, val)
		this.arr = append(this.arr, val)
		this.length++
		return
	} 
	
	if this.mins[this.length-1] >= val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[this.length-1])
	}
	this.arr = append(this.arr, val)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	this.arr = this.arr[:this.length-1]
	this.mins = this.mins[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.arr[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.length-1]
}
