/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
	resmap := make(map[int][]int, 0)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int){
		if root == nil {
			return
		}
		resmap[depth] = append(resmap[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	fmt.Println(resmap)
	for i := 0; i< len(resmap); i++ {
		res = append(res, resmap[i])
	}

	return res
}
