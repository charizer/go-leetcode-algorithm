package valid_parentheses

//https://leetcode-cn.com/problems/valid-parentheses/

func IsValid(s string) bool {
	leftx, rightx, leftz, rightz, leftd, rightd := '(', ')', '[', ']', '{', '}'
	unMatch := []int32{}
	for idx, c := range s {
		if idx == 0 {
			unMatch = append(unMatch, c)
			continue
		}
		if c == leftx || c == leftz || c == leftd {
			unMatch = append(unMatch, c)
		}else{
			if len(unMatch) <= 0 {
				return false
			}
			if c == rightx && unMatch[len(unMatch)-1] == leftx ||
				c == rightz && unMatch[len(unMatch)-1] == leftz ||
				c == rightd && unMatch[len(unMatch)-1] == leftd {
				unMatch = unMatch[0:len(unMatch)-1]
				continue
			}
			return false
		}
	}
	if len(unMatch) <= 0 {
		return true
	}else{
		return false
	}
}


func IsValid2(s string) bool {
	n := len(s)
	// 如果字符串长度是奇数，肯定不匹配
	if n % 2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// 左括号，压栈，栈中暂存对应的右括号
		if s[i] == '[' {
			stack = append(stack, ']')
		}else if s[i] == '(' {
			stack = append(stack, ')')
		}else if s[i] == '{' {
			stack = append(stack, '}')
		}else{
			// 右括号，从栈中弹出对应的括号
			// 不相等，则不匹配，直接返回
			if len(stack) <=0 || stack[len(stack)-1] != s[i] {
				return false
			}else{
				// 否则弹出相应括号，继续比较
				stack = stack[:len(stack)-1]
			}
		}
	}
	// 括号匹配完则符合条件
	return len(stack) == 0
}