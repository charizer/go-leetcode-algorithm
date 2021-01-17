package binary_tree_postorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PostorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, PostorderTraversal2(root.Left)...)
	ans = append(ans, PostorderTraversal2(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

func PostorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			// 将数据插入头部
			ans = append([]int{p.Val}, ans...)
			// 把根节点放入栈中
			stack = append(stack, p)
			// 处理右子树
			p = p.Right
		}else{
			// 弹出节点
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 处理左子树
			p = node.Left
		}
	}
	return ans
}

func PostorderTraversal3(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, PostorderTraversal3(root.Left)...)
	ans = append(ans, PostorderTraversal3(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

// 前序： 左 根 右
// 后续： 左 右 根
// 中序： 根 左 右   --- 根，右，左反转

func PostorderTraversal4(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			ans = append([]int{p.Val}, ans...)
			stack = append(stack, p)
			p = p.Right
		}else{
			p = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			p = p.Left
		}
	}
	return ans
}

