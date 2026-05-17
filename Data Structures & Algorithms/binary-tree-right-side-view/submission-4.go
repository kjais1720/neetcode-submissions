/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	// curr := &TreeNode{
	// 	Right: root,
	// }
	dfs(root, 0)
	// for i := 0; i<len(mp); i++ {
	// 	res = append(res, mp[i])
	// }
	return res
}
