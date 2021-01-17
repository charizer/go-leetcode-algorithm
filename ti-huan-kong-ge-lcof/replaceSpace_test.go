package ti_huan_kong_ge_lcof

import "testing"

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	want := "We%20are%20happy."
	got := ReplaceSpace(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
