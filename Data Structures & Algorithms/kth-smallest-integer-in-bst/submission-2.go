/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func itemsInOrder (root *TreeNode, mp map[int]int) {
	if root == nil {
		return 
	}

	itemsInOrder(root.Left, mp)
	mp[len(mp)+1] = root.Val
	itemsInOrder(root.Right, mp)
	return
}

func kthSmallest(root *TreeNode, k int) int {
    mp := make(map[int]int, 0)
	itemsInOrder(root, mp)
	return mp[k]
}
