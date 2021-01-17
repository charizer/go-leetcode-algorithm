package lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	want := 3
	got := LengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "bbbbb"
	want = 1
	got = LengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "pwwkew"
	want = 3
	got = LengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	s := "abcabcbb"
	want := 3
	got := LengthOfLongestSubstring2(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "bbbbb"
	want = 1
	got = LengthOfLongestSubstring2(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "pwwkew"
	want = 3
	got = LengthOfLongestSubstring2(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
