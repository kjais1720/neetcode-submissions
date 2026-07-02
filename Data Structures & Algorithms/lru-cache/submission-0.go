type Node struct {
	prev *Node
	val int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) push(n int) {
	if l.head == nil {
		l.head = &Node{
			val: n,
		}
		l.tail = l.head
		return
	}

	l.tail.next = &Node{
		prev: l.tail,
		val: n,
	}
	l.tail = l.tail.next
}

func (l *List) pop(n *Node) {
	if n == nil {
		return
	}
	if n == l.head {
		l.head = n.next
	}
	if n == l.tail {
		l.tail = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

type cacheObj struct {
	val int
	pos *Node
}

type LRUCache struct {
    cache map[int]cacheObj
	capacity int
	list *List
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		cache: make(map[int]cacheObj, capacity),
		capacity: capacity,
		list: &List{},
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if ok && this.list.tail.val != key { // if its not at tail, then repostion it
		this.list.pop(v.pos)
		delete(this.cache, key)
		this.Put(key, v.val)
	} else if !ok {
		return -1
	}
	return v.val
	
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if v, ok := this.cache[key]; ok {
		this.list.pop(v.pos)
		delete(this.cache, key)
	}

	if len(this.cache) == this.capacity {
		delete(this.cache, this.list.head.val)
		this.list.pop(this.list.head)
	}
	this.list.push(key)
	this.cache[key] = cacheObj{
		val: value,
		pos: this.list.tail,
	}

}
