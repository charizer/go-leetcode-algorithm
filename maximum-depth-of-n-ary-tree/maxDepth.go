package maximum_depth_of_n_ary_tree

type Node struct {
	Val int
	Children []*Node
}

func MaxDepth(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	max := 1
	for _, child := range root.Children {
		cur := MaxDepth(child) + 1
		if  cur > max {
			max = cur
		}
	}
	return max
}

func MaxDepth2(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		ans++
		size := len(queue)
		for i := 0 ; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
	}
	return ans
}
