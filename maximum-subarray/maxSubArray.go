package maximum_subarray

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur > 0 {
			cur += nums[i]
		}else{
			cur = nums[i]
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
