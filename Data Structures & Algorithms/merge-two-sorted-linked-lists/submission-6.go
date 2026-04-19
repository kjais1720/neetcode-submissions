/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
	var tail *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil && list2 != nil {
			if head == nil {
				head = list2
				tail = list2
			} else {
				tail.Next = list2
				tail = list2
			}
			list2 = list2.Next
		} else if list2 == nil && list1 != nil {
			if head == nil {
				head = list1
				tail = list1
			} else {
				tail.Next = list1
				tail = list1
			}
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			if head == nil {
				head = list1
				tail = list1
			} else {
				tail.Next = list1
				tail = list1
			}
			list1 = list1.Next
		} else if list2.Val <= list1.Val {
			if head == nil {
				head = list2
				tail = list2
			} else {
				tail.Next = list2
				tail = list2
			}
			list2 = list2.Next
		}
	}
	return head
}
