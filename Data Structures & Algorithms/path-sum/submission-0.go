/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	var hasLeaf func(root *TreeNode, sum int) bool
	hasLeaf = func(root *TreeNode, sum int) bool{
		if root == nil {
			return false
		}
		sum += root.Val
		if sum == targetSum && root.Left == nil && root.Right == nil {
			return true
		} else if hasLeaf(root.Left, sum) {
			return true
		} else if hasLeaf(root.Right, sum) {
			return true
		} else {
			return false
		}

	}

	return hasLeaf(root, 0)
}
