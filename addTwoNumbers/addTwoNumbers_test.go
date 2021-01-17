package addTwoNumbers

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

func TestAddTwoNumbers(t *testing.T) {
	l1 := createList([]int{2,4,3})
	l2 := createList([]int{5,6,4})
	want := []int{7,0,8}
	got := AddTwoNumbers(l1,l2)
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		idx++
		got = got.Next
	}
	l1 = createList([]int{1,9})
	l2 = createList([]int{9})
	want = []int{0,0,1}
	got = AddTwoNumbers(l1,l2)
	idx = 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		idx++
		got = got.Next
	}
}