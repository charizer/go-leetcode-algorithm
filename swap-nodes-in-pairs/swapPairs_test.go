package swap_nodes_in_pairs

import "testing"

func createList(arr []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, item := range arr {
		pre.Next = &ListNode{Val: item}
		pre = pre.Next
	}
	return head.Next
}

func TestSwapPairs(t *testing.T) {
	l1 := createList([]int{1,2,3,4})
	want := []int{2,1,4,3}
	got := SwapPairs(l1)
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		idx++
		got = got.Next
	}
}

func TestSwapPairs2(t *testing.T) {
	l1 := createList([]int{1,2,3,4})
	want := []int{2,1,4,3}
	got := SwapPairs2(l1)
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		idx++
		got = got.Next
	}
}
