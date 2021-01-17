package cong_wei_dao_tou_da_yin_lian_biao_lcof

type ListNode struct {
	Val int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(ReversePrint(head.Next), head.Val)
}
