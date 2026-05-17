/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    mp := make(map[int][]int)

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		mp[depth] = append(mp[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	res := make([]int, 0, len(mp))
	for i := 0; i<len(mp); i++ {
		res = append(res, mp[i][len(mp[i])-1])
	}
	return res
}
