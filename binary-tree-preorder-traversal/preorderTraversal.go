package binary_tree_preorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PreorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, PreorderTraversal2(root.Left)...)
	ans = append(ans, PreorderTraversal2(root.Right)...)
	return ans
}

func PreorderTraversal1(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 先将根节点加入结果集
			ans = append(ans, root.Val)
			// 将遍历过的结点加入栈，后序出栈依次访问结点的右子树
			stack = append(stack, root)
			// 遍历当前结点的左子树
			root = root.Left
		}
		// 节点左子树上边已经处理完成，接着处理右子树
		root = stack[len(stack)-1].Right
		stack = stack[0:len(stack)-1]
	}
	return ans
}

func PreorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			// 取根节点的值
			ans = append(ans, p.Val)
			// 把根节点放入栈中
			stack = append(stack, p)
			// 处理左子树
			p = p.Left
		}else{
			// 弹出节点
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 处理右子树
			p = node.Right
		}
	}
	return ans
}



func PreorderTraversal3(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			ans = append(ans, p.Val)
			// 用于后边弹出时处理右子树
			stack = append(stack, p)
			p = p.Left
		}else{
			// 弹出节点处理右子树
			p = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			p = p.Right
		}
	}
	return ans
}

func PreorderTraversal5(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, PreorderTraversal5(root.Left)...)
	ans = append(ans, PreorderTraversal5(root.Right)...)
	return ans
}

