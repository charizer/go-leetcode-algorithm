package maxSlidingWindow

import "math"

//暴力求解
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	result := []int{}
	for i := 0; i<= len(nums) - k ; i++ {
		result = append(result, max(nums[i:i+k]))
	}
	return result
}

func max(nums []int) int {
	m := math.MinInt32
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

// https://leetcode-cn.com/problems/sliding-window-maximum/
// https://www.bilibili.com/video/BV1rE411L77c?from=search&seid=11558344942686276789
func MaxSlidingWindow2(nums []int, k int) []int {
	result := []int{}
	if k <= 0 || len(nums) < k {
		return result
	}
	//存放元素的位置
	queue := []int{}
	for i := 0; i < len(nums); i++ {
		//删掉队列中比当前元素nums[i]小的所有元素
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[0:len(queue)-1]
		}
		//将当前元素的位置放入队列
		queue = append(queue,i)
		//清除队列中超出窗口范围的元素
		for len(queue) > 0 && queue[0] < i - k + 1 {
			queue = queue[1:]
		}
		//如果窗口里元素数量已经有K个
		if i >= k - 1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}