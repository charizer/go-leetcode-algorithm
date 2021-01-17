package valid_anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		sm[s[i]]++
		tm[t[i]]++
	}
	if len(sm) != len(tm) {
		return false
	}
	for k, _ := range sm {
		if _, ok := tm[k];ok {
			if tm[k] != sm[k] {
				return false
			}
		}else{
			return false
		}
	}
	return true
}
