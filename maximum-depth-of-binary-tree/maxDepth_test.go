package maximum_depth_of_binary_tree

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

func TestMaxDepth(t *testing.T) {
	root := createTree()
	want := 3
	got := MaxDepth(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMaxDepth2(t *testing.T) {
	root := createTree()
	want := 3
	got := MaxDepth2(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMaxDepth3(t *testing.T) {
	root := createTree()
	want := 3
	got := MaxDepth3(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMaxDepth4(t *testing.T) {
	root := createTree()
	want := 3
	got := MaxDepth4(root)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
