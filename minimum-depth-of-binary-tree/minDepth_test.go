package minimum_depth_of_binary_tree

import "testing"

func createTree() *TreeNode{
	root := &TreeNode{Val: 3}
	lc := &TreeNode{Val: 9}
	rc := &TreeNode{Val: 20}
	rcl := &TreeNode{Val: 15}
	rcr := &TreeNode{Val:7}
	root.Left = lc
	root.Right = rc
	rc.Left = rcl
	rc.Right = rcr
	return root
}

func TestMinDepth(t *testing.T) {
	root := createTree()
	want := 2
	got := MinDepth(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMinDepth2(t *testing.T) {
	root := createTree()
	want := 2
	got := MinDepth2(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMinDepth4(t *testing.T) {
	root := createTree()
	want := 2
	got := MinDepth4(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
