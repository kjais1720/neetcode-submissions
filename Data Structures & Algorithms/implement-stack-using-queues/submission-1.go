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
	q MyQ
}

func Constructor() MyStack {
	return MyStack{
		q: NewQ(),
	}
}

func (this *MyStack) Push(x int) {
	i := 0
	this.q.Enq(x)
	for i < this.q.size - 1 {
		n := this.q.Deq()
		this.q.Enq(n)
		i++
	}

}

func (this *MyStack) Pop() int {
	return this.q.Deq()
}

func (this *MyStack) Top() int {
	return this.q.Peek()
}

func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
