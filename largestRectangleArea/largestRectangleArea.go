package largestRectangleArea

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
//精解： https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/84-by-ikaruga/
func LargestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	largest := 0
	for i := 0; i < len(heights); i++ {
		l, r := getMinWidth(i, heights)
		tmp := (l + r + 1) * heights[i]
		if tmp > largest {
			largest = tmp
		}
	}
	return largest
}

func getMinWidth(idx int, height []int) (int, int) {
	left, right := 0, 0
	for i := idx-1; i >= 0; i--{
		if height[i] >= height[idx] {
			left++
		}else{
			break
		}
	}
	for i := idx+1; i < len(height); i++{
		if height[idx] <= height[i] {
			right++
		}else{
			break
		}
	}
	return left, right
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LargestRectangleArea2(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	heights = append([]int{0}, heights...)
	heights = append(heights,0)
	largest := 0
	stack := []int{}
	for i:=0; i<len(heights);i++ {
		for len(stack) > 0 &&  heights[i] < heights[stack[len(stack)-1]] {
			cur := stack[len(stack)-1]
			right := i -1
			stack = stack[0:len(stack)-1]
			left := stack[len(stack)-1] + 1
			largest = max(largest, (right-left+1)*heights[cur])
		}
		stack = append(stack,i)
	}
	return largest
}
