package positions_of_large_groups

import "testing"

func TestLargeGroupPositions(t *testing.T) {
	s := "abbxxxxzzy"
	want := [][]int{{3, 6}}
	got := LargeGroupPositions(s)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], gg)
			}
		}
	}

	s = "abc"
	want = [][]int{{}}
	got = LargeGroupPositions3(s)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], gg)
			}
		}
	}

	s = "abcdddeeeeaabbbcd"
	want = [][]int{{3,5},{6,9},{12,14}}
	got = LargeGroupPositions3(s)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], gg)
			}
		}
	}

	s = "aaa"
	want = [][]int{{0,2}}
	got = LargeGroupPositions3(s)
	if len(want) != len(got) {
		t.Errorf("want:%d got:%d", len(want),len(got))
	}else{
		for idx, g := range got {
			for i, gg := range g {
				if gg != want[idx][i] {
					t.Errorf("want:%d got:%d", want[idx][i], gg)
				}
			}
		}
	}
}
