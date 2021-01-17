package trapping_rain_water

func max(a, b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}
func Trap(height []int) int {
	if len(height) <= 0 {
		return 0
	}
	n := len(height)
	left, right, leftMax, rightMax, total := 0, n-1, height[0], height[n-1], 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			total += leftMax - height[left]
			left++
		}else{
			total += rightMax - height[right]
			right--
		}
	}
	return total
}
