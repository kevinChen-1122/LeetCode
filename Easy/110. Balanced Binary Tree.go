package Easy

// https://leetcode.com/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)
	diff := abs(leftHeight - rightHeight)
	if isLeftBalanced && isRightBalanced && diff <= 1 {
		return true, 1 + max(leftHeight, rightHeight)
	}
	return false, -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
