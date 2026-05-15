/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type qu struct {
	arr []*TreeNode
}

func (q *qu) enq(n *TreeNode) {
	q.arr = append(q.arr, n)
}

func (q *qu) deq() *TreeNode {
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
} 

func levelOrder(root *TreeNode) [][]int {
    q := &qu{}
	if root == nil {
		return [][]int{}
	}
	q.enq(root)
	res := make([][]int, 0, 1000)
	for len(q.arr) > 0 {
		lenq := len(q.arr)
		arr := make([]int, 0, lenq)
		for i := 0; i < lenq; i++ {
			n := q.deq()
			arr = append(arr, n.Val)
			if n.Left != nil {q.enq(n.Left)}
			if n.Right != nil {q.enq(n.Right)}
		}
		res = append(res, arr)
	}
	return res
}
