package cong_wei_dao_tou_da_yin_lian_biao_lcof

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

func TestReversePrint(t *testing.T) {
	node := createList([]int{1,3,2})
	want := []int{2,3,1}
	got := ReversePrint(node)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
}
