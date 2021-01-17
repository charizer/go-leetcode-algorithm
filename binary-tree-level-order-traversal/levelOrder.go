package binary_tree_level_order_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LevelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	// 将根节点放入队列中，然后不断遍历队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 获取当前队列的长度，这个长度相当于当前这一层的节点个数
		levelNum := len(queue)
		// 将队列中的元素都拿出来(也就是获取这一层的节点)，放到临时list中
		sub := []int{}
		for i := 0; i < levelNum; i++ {
			// 队列中元素出队，即依次取出一层的元素
			node := queue[0]
			queue = queue[1:]
			// 下一层元素入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// 记录当前层元素
			sub = append(sub, node.Val)
		}
		// 记录一层的元素
		ans = append(ans, sub)
	}
	return ans
}

func LevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	dfs(&ans, root, 0)
	return ans
}

func dfs (ans *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	// ans 中还没有创建 level 对应的列表时
	// 应该在 ans 中新建一个列表用来保存该 level 的所有节点。
	if level >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	// 将节点加入当前层
	(*ans)[level] = append((*ans)[level], root.Val)
	dfs(ans, root.Left, level + 1)
	dfs(ans, root.Right, level + 1)
}

//bfs
func LevelOrder3(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	p := root
	queue := []*TreeNode{p}
	for len(queue) > 0 {
		sub := []int{}
		size := len(queue)
		for i :=0 ; i < size; i++ {
			p = queue[0]
			queue = queue[1:]
			sub = append(sub, p.Val)
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		ans = append(ans, sub)
	}
	return ans
}

// dfs
func LevelOrder4(root *TreeNode) [][]int {
	ans := [][]int{}
	dfs4(root, 0, &ans)
	return ans
}

func dfs4(root *TreeNode, level int, ans *[][]int) {
	if root == nil {
		return
	}
	if level >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], root.Val)
	dfs4(root.Left, level+1, ans)
	dfs4(root.Right, level+1, ans)
}


