package group_anagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	tmp := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})
		tmp[string(s)] = append(tmp[string(s)], str)
	}
	for _, strs := range tmp {
		t := []string{}
		for _, str := range strs {
			t = append(t, str)
		}
		ans = append(ans, t)
	}
	return ans
}

func GroupAnagrams2(strs []string) [][]string {
	ans := [][]string{}
	tmp := make(map[[26]int][]string)
	for _, str := range strs {
		chs := [26]int{}
		for _, ch := range str {
			chs[ch-'a']++
		}
		tmp[chs] = append(tmp[chs], str)
	}
	for _, v := range tmp {
		ans = append(ans, v)
	}
	return ans
}
