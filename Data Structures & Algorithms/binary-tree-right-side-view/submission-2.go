/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    mp := make(map[int]int)

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if _, ok := mp[depth]; !ok {
			mp[depth] = root.Val
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	// curr := &TreeNode{
	// 	Right: root,
	// }
	dfs(root, 0)
	res := make([]int, 0, len(mp))
	for i := 0; i<len(mp); i++ {
		res = append(res, mp[i])
	}
	fmt.Println(mp)
	fmt.Println(res)
	return res
}
