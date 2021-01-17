package waysToStep

import "testing"

func TestWaysToStep(t *testing.T) {
	n := 3
	want := 4
	got := WaysToStep(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 5
	want = 13
	got = WaysToStep(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestWaysToStep2(t *testing.T) {
	n := 3
	want := 4
	got := WaysToStep2(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 5
	want = 13
	got = WaysToStep2(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}