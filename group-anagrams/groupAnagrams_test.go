package group_anagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	want := [][]string{{"eat","tea","ate"}, {"tan","nat"}, {"bat"}}
	got := GroupAnagrams2(strs)
	for idx, g := range got {
		for i, gg := range g {
			if gg != want[idx][i] {
				t.Errorf("want:%s got:%s", want[idx][i], gg)
			}
		}
	}
}
