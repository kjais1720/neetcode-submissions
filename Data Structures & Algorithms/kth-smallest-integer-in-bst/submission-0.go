/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func itemsInOrder (root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := append(itemsInOrder(root.Left), root.Val)
	res = append(res, itemsInOrder(root.Right)...)
	return res
}

func kthSmallest(root *TreeNode, k int) int {
    return itemsInOrder(root)[k-1]
	
}
