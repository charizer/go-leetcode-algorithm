package last_stone_weight

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push (x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

// 一定要用heap的Push和Pop方法，heap会调整，千万不能用自己的Pop和Push方法
func LastStoneWeight3(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for len(h) > 1 {
		l := heap.Pop(&h)
		r := heap.Pop(&h)
		left := l.(int) - r.(int)
		if left > 0 {
			heap.Push(&h, left)
		}
	}
	if len(h) > 0 {
		return h[0]
	}
	return 0
}

// 暴力求解
func LastStoneWeight(stones []int) int {
	// 一块石头都没有，直接返回0
	if len(stones) <= 0 {
		return 0
	}
	// 还有多余两块石头，则进行处理
	for len(stones) > 1 {
		// 先排序
		sort.Ints(stones)
		n := len(stones)
		// 处理最大的两块
		left := stones[n-1] - stones[n-2]
		// 把最大两块从数组的去掉
		stones = stones[0:n-2]
		// 如果这两块有剩余，则将剩余入数组
		if left != 0 {
			stones = append(stones, left)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}else{
		return 0
	}
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   {
	a := h.IntSlice
	if len(a) <= 0 {
		return nil
	}
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func LastStoneWeight2(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
