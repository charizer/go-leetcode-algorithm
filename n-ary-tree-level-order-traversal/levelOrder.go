package n_ary_tree_level_order_traversal

type Node struct {
	Val int
	Children []*Node
}

func LevelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		sub := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sub = append(sub, node.Val)
			queue = append(queue, node.Children...)
		}
		ans = append(ans, sub)
	}
	return ans
}

func dfs(ans *[][]int, level int, root *Node) {
	if root == nil {
		return
	}
	if level >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], root.Val)
	for _, child := range root.Children {
		dfs(ans, level+1, child)
	}
}

func LevelOrder2(root *Node) [][]int {
	ans := [][]int{}
	dfs(&ans, 0, root)
	return ans
}
