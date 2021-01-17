package maximum_depth_of_binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}

//dfs
func MaxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	return max(MaxDepth(root.Left) + 1, MaxDepth(root.Right)+1)
}

//bfs
func MaxDepth2(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
	}
	return depth
}

func MaxDepth3(root *TreeNode) int {
	ans := 0
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans ++
		size := len(queue)
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
	}
	return ans
}

func MaxDepth4(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	return max(MaxDepth4(root.Left), MaxDepth4(root.Right)) + 1
}
