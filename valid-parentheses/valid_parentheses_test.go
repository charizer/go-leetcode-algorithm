package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	s := "()"
	want := true
	got := IsValid(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "()[]{}"
	want = true
	got = IsValid(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "(]"
	want = false
	got = IsValid(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "([)]"
	want = false
	got = IsValid(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}

func TestIsValid2(t *testing.T) {
	s := "()"
	want := true
	got := IsValid2(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "()[]{}"
	want = true
	got = IsValid2(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "(]"
	want = false
	got = IsValid2(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "([)]"
	want = false
	got = IsValid2(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}