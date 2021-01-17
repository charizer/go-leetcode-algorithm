package add_digits

import "testing"

func TestAddDigits(t *testing.T) {
	num := 38
	want := 2
	got := AddDigits(num)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
