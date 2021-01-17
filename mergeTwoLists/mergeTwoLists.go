package mergeTwoLists

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pre := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1//&ListNode{Val: l1.Val}
			l1 = l1.Next
		}else{
			pre.Next = l2//&ListNode{Val: l2.Val}
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
	return head.Next
}
