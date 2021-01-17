package binary_tree_level_order_traversal

import "testing"

func TestLevelOrder3(t *testing.T) {
	root := &TreeNode{Val:3}
	left1 := &TreeNode{Val: 9}
	right1 := &TreeNode{Val: 20}
	root.Left = left1
	root.Right = right1
	rl := &TreeNode{Val: 15}
	rr := &TreeNode{Val: 7}
	right1.Left = rl
	right1.Right = rr
	want := [][]int{{3},{9,20},{15,7}}
	got := LevelOrder3(root)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], gg)
			}
		}
	}
}

func TestLevelOrder4(t *testing.T) {
	root := &TreeNode{Val:3}
	left1 := &TreeNode{Val: 9}
	right1 := &TreeNode{Val: 20}
	root.Left = left1
	root.Right = right1
	rl := &TreeNode{Val: 15}
	rr := &TreeNode{Val: 7}
	right1.Left = rl
	right1.Right = rr
	want := [][]int{{3},{9,20},{15,7}}
	got := LevelOrder4(root)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], gg)
			}
		}
	}
}
