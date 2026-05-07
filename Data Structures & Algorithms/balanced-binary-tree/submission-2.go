/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func calcHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	
	rightH :=  1 + calcHeight(root.Right)

	leftH := 1 + calcHeight(root.Left)
	
	if rightH > leftH {
		return rightH
	}
	return leftH
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
    leftHeight := calcHeight(root.Left)
	rightHeight := calcHeight(root.Right)
	fmt.Println(leftHeight, rightHeight)
	if math.Abs(float64(rightHeight - leftHeight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}
