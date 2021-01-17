package swap_nodes_in_pairs

type ListNode struct {
	Val int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := pre.Next.Next
		pre.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		pre = node1
	}
	return dummy.Next
}

func SwapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = SwapPairs2(newHead.Next)
	newHead.Next = head
	return newHead
}
