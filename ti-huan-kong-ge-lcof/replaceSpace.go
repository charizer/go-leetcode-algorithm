package ti_huan_kong_ge_lcof

func ReplaceSpace(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans += "%20"
		}else{
			ans += string(s[i])
		}
	}
	return ans
}
