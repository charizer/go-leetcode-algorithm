package sliding_window_maximum

func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result := []int{}
	if k <= 0 || n < k {
		return result
	}
	// 存放元素的小标，用于清理队列中超过范围的数据
	dqueue := []int{}
	for i:= 0; i < n; i++ {
		// 从队列中弹出比当前元素小的元素， 从尾部弹出
		for len(dqueue) > 0 && nums[i] > nums[dqueue[len(dqueue)-1]] {
			dqueue = dqueue[0:len(dqueue)-1]
		}
		// 将当前元素入队
		dqueue = append(dqueue, i)
		// 将超出k窗口的元素弹出， 从头部弹出, 这里使用到小标判断是否在k范围内
		for len(dqueue) > 0 && dqueue[0] < i - k + 1 {
			dqueue = dqueue[1:]
		}
		// 达到K个元素才输出数据， 队头为窗口内最大值
		if i >= k-1 && len(dqueue) > 0 {
			result = append(result, nums[dqueue[0]])
		}
	}
	return result
}
