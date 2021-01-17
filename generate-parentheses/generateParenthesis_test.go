package generate_parentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	want := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	got := GenerateParenthesis(n)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%s got:%s", want[idx], got[idx])
		}
	}
}
