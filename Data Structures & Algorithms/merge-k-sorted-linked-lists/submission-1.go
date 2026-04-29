/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	var newHead *ListNode
	var newTail *ListNode
	
	nonEmptyLists := len(lists)

	// [1,1,3]
	// [2,1,3]
	// [2,3,3]
	//

	for nonEmptyLists > 1 {
		nonEmptyLists = len(lists)
		min :=  1001
		smallestNodeIdx := 0
		for idx, node := range lists {
			if node == nil {
				nonEmptyLists--
				continue
			}
			if node.Val < min {
				min = node.Val
				smallestNodeIdx = idx
			}
		}
		lists[smallestNodeIdx] = lists[smallestNodeIdx].Next

		if newHead == nil {
			newHead = &ListNode{
				Val: min,
			}
			newTail = newHead
			continue
		}
		// fmt.Println("head", newHead.Val)
		// fmt.Println("tail", newTail.Val)
		newTail.Next = &ListNode{
			Val: min,
		}
		newTail = newTail.Next
		// arr := []int{}
		// tempHead := newHead
		// for tempHead != nil {
		// 	arr = append(arr, tempHead.Val)
		// 	tempHead = tempHead.Next
		// }
		// fmt.Println("arr",arr)
	}
	for _, node := range lists {
		if node != nil {
			newTail.Next = node
			break
		}
	}
	return newHead
}
