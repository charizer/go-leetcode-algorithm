package min_stack

import "testing"

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	want := -3
	got := minStack.GetMin()
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	minStack.Pop()
	want = 0
	got = minStack.Top()
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	want = -2
	got = minStack.GetMin()
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
