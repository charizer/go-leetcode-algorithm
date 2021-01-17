package lowest_common_ancestor_of_a_binary_search_tree

func createTree() *TreeNode {
	root := &TreeNode{Val: 6}
	l1 := &TreeNode{Val: 2}
	r1 := &TreeNode{Val: 8}
	root.Left = l1
	root.Right = r1
	l1l := &TreeNode{Val: 0}
	l1r := &TreeNode{Val: 4}
	l1.Left = l1l
	l1.Right = l1r
	l2l := &TreeNode{Val: 3}
	l2r := &TreeNode{Val: 5}
	l1r.Left = l2l
	l1r.Right = l2r
	r1l := &TreeNode{Val: 7}
	r1r := &TreeNode{Val: 9}
	r1.Left = r1l
	r1.Right = r1r
	return root
}
