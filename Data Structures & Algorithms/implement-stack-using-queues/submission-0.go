type MyQ struct {
	arr []int
	size int
}

func NewQ() MyQ {
	return MyQ{
		arr: []int{}, 
	}
}

func (this *MyQ) Enq(x int) {
	this.arr = append(this.arr, x)
	this.size++
}

func (this *MyQ) Deq() int {
	if this.IsEmpty() {
		return -1
	}
	res := this.arr[0]
	this.arr = this.arr[1:this.size]
	this.size--
	return res
}

func (this *MyQ) Peek() int {
	return this.arr[0]
}

func (this *MyQ) IsEmpty() bool {
	return this.size == 0
}

type MyStack struct {
	q1 MyQ
	q2 MyQ
}

func Constructor() MyStack {
	return MyStack{
		q1: NewQ(),
		q2: NewQ(),

	}
}

func (this *MyStack) Push(x int) {
	if this.q1.IsEmpty() {
		this.q1.Enq(x)
		for !this.q2.IsEmpty() {
			this.q1.Enq(this.q2.Deq())
		}
	} else if this.q2.IsEmpty() {
		this.q2.Enq(x)
		for !this.q1.IsEmpty() {
			this.q2.Enq(this.q1.Deq())
		}
	}
}

func (this *MyStack) Pop() int {
	if !this.q1.IsEmpty() {
		return this.q1.Deq()
	} else if !this.q2.IsEmpty() {
		return this.q2.Deq()
	} else {
		return -1
	}
	return -1
}

func (this *MyStack) Top() int {
	if !this.q1.IsEmpty() {
		return this.q1.Peek()
	} else if !this.q2.IsEmpty() {
		return this.q2.Peek()
	} else {
		return -1
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return this.q1.IsEmpty() && this.q2.IsEmpty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
