package n_ary_tree_postorder_traversal

type Node struct {
	Val int
	Children []*Node
}

// 递归法
func Postorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	for _, child := range root.Children {
		ans = append(ans, Postorder(child)...)
	}
	ans = append(ans, root.Val)
	return ans
}

// 迭代法
func Postorder2(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		ans = append([]int{node.Val}, ans...)
		for _, child := range node.Children {
			stack = append(stack, child)
		}
	}
	return ans
}
