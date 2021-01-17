package merge_two_sorted_lists

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	pre := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}else{
			pre.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return newHead.Next
}
