package fibonacci_number

import "testing"

func TestFib(t *testing.T) {
	n := 2
	want := 1
	got := Fib(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 3
	want = 2
	got = Fib(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 4
	want = 3
	got = Fib(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
