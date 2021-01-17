package generate_parentheses

func GenerateParenthesis(n int) []string {
	ans := []string{}
	if n <= 0 {
		return ans
	}
	generate(&ans, "",0, 0, n)
	return ans
}

func generate(ans* []string, s string, left, right int, n int) {
	if left == n && right == n {
		*ans = append(*ans, s)
		return
	}
	if left < n {
		generate(ans, s + "(", left+1, right, n)
	}
	if right < left {
		generate(ans, s + ")", left, right+1, n)
	}
}
