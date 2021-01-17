package binary_tree_inorder_traversal

import "testing"

func TestInorderTraversal4(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{1,3,2}
	got := InorderTraversal4(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}

func TestInorderTraversal5(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{1,3,2}
	got := InorderTraversal5(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}
