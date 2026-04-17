type LinkedList struct {
    Head *Node
    Tail *Node
}

type Node struct {
    Val int
    Next *Node
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
    node := ll.Head
    for i := 0; i <= index; i++ {
        if node == nil {
            return -1
        }
        if i == index {
            return node.Val
        }
        node = node.Next
    }
    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := &Node{Val: val, Next: ll.Head}
    ll.Head = newNode
    if ll.Tail == nil {
        ll.Tail = newNode
    }
}

func (ll *LinkedList) InsertTail(val int) {
    newNode := &Node{Val: val}
    if ll.Head == nil {
        ll.Head = newNode
        ll.Tail = newNode
        return
    }
    ll.Tail.Next = newNode
    ll.Tail = newNode
}

func (ll *LinkedList) Remove(index int) bool {
    if ll.Head == nil {
        return false
    }
    if index == 0 {
        if ll.Head == ll.Tail {
            ll.Tail = nil
        }
        ll.Head = ll.Head.Next
        return true
    }
    prev := ll.Head
    for i := 0; i < index-1; i++ {
        if prev.Next == nil {
            return false
        }
        prev = prev.Next
    }
    if prev.Next == nil {
        return false
    }
    if prev.Next == ll.Tail {
        ll.Tail = prev
    }
    prev.Next = prev.Next.Next
    return true
}

func (ll *LinkedList) GetValues() []int {
    res := make([]int, 0)
    node := ll.Head
    for node != nil {
        res = append(res, node.Val)
        node = node.Next
    }
    return res
}