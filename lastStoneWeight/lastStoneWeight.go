package lastStoneWeight

import "sort"

//https://leetcode-cn.com/problems/last-stone-weight/
func LastStoneWeight(stones []int) int {
	if len(stones) <= 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	for len(stones) >= 2 {
		sort.Ints(stones)
		n := len(stones)
		left := stones[n-1] - stones[n-2]
		stones = stones[0:n-2]
		if left > 0 {
			stones = append(stones, left)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}else{
		return 0
	}
}
