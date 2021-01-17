package reverse_linked_list

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

func TestReverseList(t *testing.T) {
	h := createList([]int{1,2,3,4,5})
	want := []int{5,4,3,2,1}
	got := ReverseList(h)
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		got = got.Next
		idx++
	}
}

func TestReverseList2(t *testing.T) {
	h := createList([]int{1,2,3,4,5})
	want := []int{5,4,3,2,1}
	got := ReverseList2(h)
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		got = got.Next
		idx++
	}
}
