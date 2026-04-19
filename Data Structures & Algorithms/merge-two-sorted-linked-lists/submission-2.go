func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	head := list2
	var pos *ListNode
	for list1 != nil {
		node := &ListNode{
			Val: list1.Val,
		}

		if head.Val > node.Val {
			node.Next = head
			head = node
			pos = node
			list1 = list1.Next
			continue
		}
		curr := pos
		if curr == nil {
			curr = list2
		}
		next := curr.Next
		for next != nil {
			if next.Val > node.Val {
				curr.Next = node
				node.Next = next
				pos = node
				break
			}
			curr = next
			next = next.Next
		} 
		if node.Next == nil {
			curr.Next = node
		}
		
		list1 = list1.Next
	}

	return head
}
