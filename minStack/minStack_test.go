package minStack

import "testing"

func TestConstructor(t *testing.T) {
	c := Constructor()
	c.Push(-2)
	c.Push(0)
	c.Push(-3)
	got := c.GetMin()
	want := -3
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	c.Pop()
	c.Pop()
	got = c.GetMin()
	want = -2
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
