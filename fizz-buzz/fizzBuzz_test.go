package fizz_buzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	n := 15
	got := FizzBuzz(n)
	want := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%s got:%s", want[idx], g)
		}
	}
}
