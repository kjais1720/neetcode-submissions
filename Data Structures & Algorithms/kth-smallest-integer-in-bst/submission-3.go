/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func itemsInOrder (root *TreeNode, mp map[int]int, k int) {
	if root == nil {
		return 
	}

	if k == 0 {
		return
	}

	itemsInOrder(root.Left, mp, k)
	mp[len(mp)+1] = root.Val
	if len(mp) == k {
		return
	}
	itemsInOrder(root.Right, mp, k)
	return
}

func kthSmallest(root *TreeNode, k int) int {
    mp := make(map[int]int, 0)
	itemsInOrder(root, mp, k)
	return mp[k]
}
