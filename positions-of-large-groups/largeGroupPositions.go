package positions_of_large_groups

func LargeGroupPositions(s string) [][]int {
	result := [][]int{}
	if len(s) < 3 {
		return result
	}
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] != s[left] {
			if i - left  >= 3 {
				result = append(result, []int{left, i-1})
			}
			left = i
		}
	}
	if len(s) - left >= 3 {
		result = append(result, []int{left, len(s)-1})
	}
	return result
}

func LargeGroupPositions2(s string) (ans [][]int) {
	cnt := 1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				ans = append(ans, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return
}

func LargeGroupPositions3(s string) (ans [][]int) {
	left := 0
	for right := 0;  right < len(s); left = right {
		for ; right < len(s) && s[left] == s[right]; right++ {}
		if right - left >= 3 {
			ans = append(ans, []int{left, right-1})
		}
	}
	return
}
