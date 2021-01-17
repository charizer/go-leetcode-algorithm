package threeSum

import "sort"

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		target := -1 * nums[i]
		f, s := i+1, n-1
		for ;f < s; f++ {
			if f > i + 1 && nums[f] == nums[f - 1] {
				continue
			}
			for f < s && nums[f] + nums[s] > target {
				s--
			}
			if f == s {
				break
			}
			if nums[f] + nums[s] == target {
				result = append(result, []int{nums[i],nums[f],nums[s]})
			}
		}
	}
	return result
}
//[-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0]
func ThreeSum2(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	n := len(nums)
	for left := 0; left < n - 2; left++ {
		if left != 0 && nums[left] == nums[left-1] {
			continue
		}
		target := 0 - nums[left]
		mid := left + 1
		last := n - 1
		for mid < last {
			if nums[mid] + nums[last] > target {
				last--
			}else if nums[mid] + nums[last] < target {
				mid++
			}else{
				ans = append(ans, []int{nums[left], nums[mid], nums[last]})
				for mid < last && nums[mid] == nums[mid+1] {
					mid++
				}
				for mid < last && nums[last] == nums[last-1] {
					last--
				}
				last--
				mid++
			}
 		}
	}
	return ans
}


func ThreeSum3(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	n := len(nums)
	for first := 0; first < n - 2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target, second, third := 0 - nums[first], first + 1, n - 1
		for second < third {
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			}else if nums[second] + nums[third] < target {
				second++
			}else{
				third--
			}
		}
	}
	return ans
}
