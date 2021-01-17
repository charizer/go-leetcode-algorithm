package powx_n

import "testing"

func TestMyPow(t *testing.T) {
	x, n := 2.00000, 10
	want := 1024.00000
	got := MyPow(x, n)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	x, n = 2.10000, 3
	want = 9.26100
	got = MyPow(x, n)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
	x, n = 2.00000, -2
	want = 0.25000
	got = MyPow(x, n)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
