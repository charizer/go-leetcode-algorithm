package minimum_depth_of_binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}else{
		return b
	}
}

func MinDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	left := MinDepth(root.Left)
	right := MinDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left,right) + 1
}

func MinDepth2(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	depth = 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]
			if p.Left == nil && p.Right == nil {
				return depth
			}
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		depth++
	}
	return depth
}

func MinDepth3(root *TreeNode) int {
	if root.Left == nil {
		return MinDepth3(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth3(root.Left) + 1
	}
	return min(MinDepth3(root.Left), MinDepth3(root.Right)) + 1
}

func MinDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]
			if p.Left == nil && p.Right == nil {
				return depth
			}
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