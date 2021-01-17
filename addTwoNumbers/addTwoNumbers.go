package addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := &ListNode{}
	pre := head
	plus := 0
	for l1 != nil || l2 != nil {
		sum := plus
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			plus = 1
		}else{
			plus = 0
		}
		pre.Next = &ListNode{Val: sum % 10}
		pre = pre.Next
	}
	if plus != 0 {
		pre.Next = &ListNode{Val: plus}
		pre = pre.Next
	}
	return head.Next
}