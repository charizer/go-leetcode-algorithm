package get_kth_magic_number_lcci

import "testing"

func TestGetKthMagicNumber(t *testing.T) {
	k := 5
	want := 9
	got := GetKthMagicNumber(k)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
