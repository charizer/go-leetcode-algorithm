package reverse_integer

import "testing"

func TestReverse(t *testing.T) {
	num := 123
	want := 321
	got := Reverse(num)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	num = -123
	want = -321
	got = Reverse(num)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}

	num1 := int32(-2147483648)
	want1 := int32(0)
	got1 := Reverse3(num1)
	if got1 != want1 {
		t.Errorf("want:%d got:%d", want, got)
	}
}
