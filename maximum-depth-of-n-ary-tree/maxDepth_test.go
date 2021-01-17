package maximum_depth_of_n_ary_tree

import "testing"

func creatTree() *Node {
	root := &Node{Val: 1}
	c1 := &Node{Val: 3}
	c2 := &Node{Val: 2}
	c3 := &Node{Val: 4}
	root.Children = []*Node{c1,c2,c3}
	cc1 := &Node{Val: 5}
	cc2 := &Node{Val: 6}
	c1.Children = []*Node{cc1,cc2}
	return root
}

func TestMaxDepth(t *testing.T) {
	node := creatTree()
	want := 3
	got := MaxDepth(node)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMaxDepth2(t *testing.T) {
	node := creatTree()
	want := 3
	got := MaxDepth2(node)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
