package summary_ranges

import (
	"strconv"
)

func SummaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return ans
	}
	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] - nums[i-1] != 1 {
			if i - 1 > left {
				ans = append(ans, strconv.Itoa(nums[left]) + "->" + strconv.Itoa(nums[i-1]))
			}else{
				ans = append(ans, strconv.Itoa(nums[left]))
			}
			left = i
		}
	}
	if left == len(nums) -1 {
		ans = append(ans, strconv.Itoa(nums[left]))
	}else{
		ans = append(ans, strconv.Itoa(nums[left]) + "->" + strconv.Itoa(nums[len(nums)-1]))
	}
	return ans
}
