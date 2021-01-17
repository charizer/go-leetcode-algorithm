package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	pre := newHead
	// 记录进位
	plus := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + plus
		plus = tmp / 10
		pre.Next = &ListNode{Val: tmp % 10}
		l1 = l1.Next
		l2 = l2.Next
		pre = pre.Next
	}
	for l1 != nil {
		tmp := l1.Val + plus
		plus = tmp / 10
		pre.Next = &ListNode{Val: tmp % 10}
		l1 = l1.Next
		pre = pre.Next
	}
	for l2 != nil {
		tmp := l2.Val + plus
		plus = tmp / 10
		pre.Next = &ListNode{Val: tmp % 10}
		l2 = l2.Next
		pre = pre.Next
	}
	if plus != 0 {
		pre.Next = &ListNode{Val: plus}
	}
	return newHead.Next
}
