package remove_outermost_parentheses

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	s := "(()())(())"
	want := "()()()"
	got := RemoveOuterParentheses(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "(()())(())(()(()))"
	want = "()()()()(())"
	got = RemoveOuterParentheses(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "()()"
	want = ""
	got = RemoveOuterParentheses(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}
}

func TestRemoveOuterParentheses2(t *testing.T) {
	s := "(()())(())"
	want := "()()()"
	got := RemoveOuterParentheses2(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "(()())(())(()(()))"
	want = "()()()()(())"
	got = RemoveOuterParentheses2(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "()()"
	want = ""
	got = RemoveOuterParentheses2(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}
}

func TestRemoveOuterParentheses3(t *testing.T) {
	s := "(()())(())"
	want := "()()()"
	got := RemoveOuterParentheses3(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "(()())(())(()(()))"
	want = "()()()()(())"
	got = RemoveOuterParentheses3(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "()()"
	want = ""
	got = RemoveOuterParentheses3(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}
}

func TestRemoveOuterParentheses5(t *testing.T) {
	s := "(()())(())"
	want := "()()()"
	got := RemoveOuterParentheses5(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "(()())(())(()(()))"
	want = "()()()()(())"
	got = RemoveOuterParentheses5(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}

	s = "()()"
	want = ""
	got = RemoveOuterParentheses5(s)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("want:%s got:%s", string(want[i]), string(got[i]))
		}
	}
}