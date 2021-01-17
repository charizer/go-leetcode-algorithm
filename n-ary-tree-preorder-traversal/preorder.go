package n_ary_tree_preorder_traversal

type Node struct {
	Val int
	Children []*Node
}

// 递归法
func Preorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	for _, child := range root.Children {
		ans = append(ans, Preorder(child)...)
	}
	return ans
}

// 迭代法
func Preorder2(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >=0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}
