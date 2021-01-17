package remove_outermost_parentheses

func RemoveOuterParentheses3(S string) string {
	now, current, ans := 0, "", ""
	for _, char := range S {
		if string(char) == "(" {
			now++
		} else if string(char) == ")" {
			now--
		}
		current += string(char)
		if now == 0 {
			ans += current[1 : len(current)-1]
			current = ""
		}
	}
	return ans
}

func RemoveOuterParentheses(S string) string {
	ans := ""
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		// 左括号，入栈
		if S[i] == '(' {
			// 加入时栈中还有括号，说明不是最外层括号
			if len(stack) > 0 {
				ans += string(S[i])
			}
			stack = append(stack, S[i])
		} else{
			// 右括号
			stack = stack[0:len(stack)-1]
			// 弹出后栈中还有元素，说明不是最外层括号
			if len(stack) > 0 {
				ans += string(S[i])
			}
		}
	}
	return ans
}

func RemoveOuterParentheses2(S string) string {
	count, ans := 0, ""
	for _, char := range S {
		if string(char) == "(" {
			if count > 0 {
				ans += string(char)
			}
			count++
		} else {
			if count > 1 {
				ans += string(char)
			}
			count--
		}
	}
	return ans
}

func RemoveOuterParentheses4(S string) string {
	if len(S) <= 0 {
		return ""
	}
	ans := ""
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		// 左括号入栈
		if S[i] == '(' {
			// 入栈是非空，则不是最外层
			if len(stack) > 0 {
				ans += string(S[i])
			}
			stack = append(stack, '(')
		}else{
			// 右括号出栈
			stack = stack[0:len(stack)-1]
			// 出栈后非空，则不是最外层
			if len(stack) > 0 {
				ans += string(S[i])
			}
		}
	}
	return ans
}

func RemoveOuterParentheses5(S string) string {
	ans := ""
	count := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' && count > 0 {
			ans += string(S[i])
		}else if S[i] == ')' && count > 1 {
			ans += string(S[i])
		}
		if S[i] == '(' {
			count++
		}else{
			count--
		}
	}
	return ans
}