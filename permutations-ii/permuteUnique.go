package permutations_ii

import "sort"

func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	ans := [][]int{}
	path := []int{}
	backtrack(nums, path, &ans, used)
	return ans
}

func backtrack (nums, path []int, ans *[][]int, used []bool){
	if len(path) == len(nums) {
		// *ans = append(*ans, path) 有问题
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack(nums, path, ans, used)
		used[i] = false
		path = path[0:len(path)-1]
	}
}
