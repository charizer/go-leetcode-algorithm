package swapPairs

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePair(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := node.Next
	node.Next = node.Next.Next
	newHead.Next = node
	return newHead
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead, tail *ListNode
	for head != nil {
		head = reversePair(head)
		if newHead == nil {
			newHead = head
		}
		if tail != nil {
			tail.Next = head
		}
		if head != nil {
			head = head.Next
			tail = head
			if head != nil {
				head = head.Next
			}
		}
	}
	return newHead
}
