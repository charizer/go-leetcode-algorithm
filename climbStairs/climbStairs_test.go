package climbStairs

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 2
	want := 2
	got := ClimbStairs2(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 3
	want = 3
	got = ClimbStairs2(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 4
	want = 5
	got = ClimbStairs2(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}