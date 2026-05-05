/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMin(root *TreeNode) int {
	curr := root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		}
		min := findMin(root.Right)
		root.Right = deleteNode(root.Right, min)
		root.Val = min
	}
	return root
}
