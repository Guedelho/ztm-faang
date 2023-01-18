package maximumdepthbinarytree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_depth := MaxDepth(root.Left)
	right_depth := MaxDepth(root.Right)
	return max(left_depth, right_depth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
