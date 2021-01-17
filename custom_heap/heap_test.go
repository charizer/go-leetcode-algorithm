package custom_heap

import (
	"testing"
)

func TestHeap_Init(t *testing.T) {
	h := Heap{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h.Init()
	t.Log(h) // [3 7 20 10 15 25 30 17 19]

	h.Push(6)
	t.Log(h) // [3 6 20 10 7 25 30 17 19 15]

	x, ok := h.Remove(5)
	t.Log(x, ok, h) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)
	t.Log(y, ok, h) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	t.Log(z, h) // 3 [7 10 15 17 19 20 30]
}
