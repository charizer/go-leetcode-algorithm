package different_ways_to_add_parentheses

import "strconv"

func DiffWaysToCompute(input string) []int {
	ans := []int{}
	for idx, ch := range input {
		if ch == '+' || ch == '-' || ch == '*' {
			left := DiffWaysToCompute(input[:idx])
			right := DiffWaysToCompute(input[idx+1:])
			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						ans = append(ans, l + r)
					case '-':
						ans = append(ans, l - r)
					case '*':
						ans = append(ans, l * r)
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		cur, _ := strconv.Atoi(input)
		ans = append(ans, cur)
	}
	return ans
}
