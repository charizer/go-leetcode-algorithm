package largest_rectangle_in_histogram

func min(a , b int) int {
	if a < b {
		return a
	}else{
		return b
	}
}

func LargestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	ans := 0
	// 左右添0，使判断边界条件统一
	heights = append([]int{0}, heights...)
	heights = append(heights,0)
	stack := []int{}
	for right := 0; right < len(heights); right++ {
		// 栈中有元素，且栈顶比当前元素小，则需要出栈，出栈的元素即为当前要计算的柱形，他前边是比他小的，后边也比他小，他就可以计算了
		for len(stack) > 0 && heights[right] < heights[stack[len(stack)-1]] {
			hmid := heights[stack[len(stack)-1]]
			// 弹出元素
			stack = stack[0:len(stack)-1]
			idxLeft := stack[len(stack)-1]
			cur := hmid * (right - idxLeft - 1)
			if cur > ans {
				ans = cur
			}
		}
		stack = append(stack, right)
	}
	return ans
}
