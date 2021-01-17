package binary_tree_postorder_traversal

import "testing"

func TestPostorderTraversal3(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{3,2,1}
	got := PostorderTraversal3(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}

func TestPostorderTraversal4(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{3,2,1}
	got := PostorderTraversal4(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}
