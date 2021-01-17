package check_if_it_is_a_straight_line

import "testing"

func TestCheckStraightLine(t *testing.T) {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	want := true
	got := CheckStraightLine(coordinates)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	coordinates = [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	want = false
	got = CheckStraightLine(coordinates)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	coordinates = [][]int{{0, 0}, {0, 1}, {0, -1}}
	want = true
	got = CheckStraightLine(coordinates)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
