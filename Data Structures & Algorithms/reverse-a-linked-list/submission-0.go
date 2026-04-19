/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	curr := &ListNode{
		Val: head.Val,
	}
	next := head.Next
	for next != nil {
		temp := next.Next
		next.Next = curr
		curr = next
		next = temp
	}
	return curr
}
