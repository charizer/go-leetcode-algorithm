package binary_tree_inorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归算法
func InorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, InorderTraversal2(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, InorderTraversal2(root.Right)...)
	return ans
}

// 迭代算法
func InorderTraversal3(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 左子树入栈，直到没有左孩子为止
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 获取到栈顶元素，再弹出，这个即为当前结果中的一个元素，可以放入ans
		root = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		ans = append(ans, root.Val)
		// 处理右子树
		root = root.Right
	}
	return ans
}

func InorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	p := root
	// p != nil, 处理左子树
	// len(stack)>0, 处理右子树
	for p != nil || len(stack) > 0 {
		if p != nil {
			// 将根节点入栈
			stack = append(stack, p)
			// 处理左子树
			p = p.Left
		}else{
			// 将根节点弹出
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 取根节点的值
			ans = append(ans, node.Val)
			// 处理右子树
			p = node.Right
		}
	}
	return ans
}

func InorderTraversal4(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		return ans
	}
	ans = append(ans, InorderTraversal4(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, InorderTraversal4(root.Right)...)
	return ans
}


func InorderTraversal5(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		}else{
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, p.Val)
			p = p.Right
		}
	}
	return ans
}

func InorderTraversal6(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, InorderTraversal6(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, InorderTraversal6(root.Right)...)
	return ans
}

func InorderTraversal7(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		}else{
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, p.Val)
			p = p.Right
		}
	}
	return ans
}