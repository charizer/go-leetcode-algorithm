package subsets

import "sort"

func Subsets(nums []int) [][]int {
	ans := [][]int{}
	items := []int{}
	//ans = append(ans, items)
	generate(0, nums, items, &ans)
	return ans
}

func generate(i int, nums []int, items []int, ans *[][]int){
	if i == len(nums) {
		*ans = append(*ans, append([]int(nil), items...))
		return
	}
	// 选择这个元素
	items = append(items, nums[i])
	generate(i+1, nums, items, ans)
	// 不选择这个元素
	items = (items)[:len(items)-1]
	generate(i+1, nums, items, ans)
}

func Subsets2(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 0 {
		return [][]int{}
	}
	items := []int{}
	ans = append(ans, items)
	backstack(0, nums, items, &ans)
	return ans
}

func backstack(level int, nums []int, items []int,ans *[][]int){
	if level == len(nums) {
		return
	}
	items = append(items, nums[level])
	*ans = append(*ans, append([]int{}, items...))
	backstack(level+1, nums, items, ans)
	items = items[:len(items)-1]
	backstack(level+1, nums, items, ans)
}

func Subsets3(nums []int) [][]int {
	ans := [][]int{}
	items := []int{}
	sort.Ints(nums)
	backtrack(nums, items, &ans, 0)
	return ans
}

func backtrack(nums, items []int, ans *[][]int, start int) {
	*ans = append(*ans, append([]int{}, items...))
	for i := start; i < len(nums); i++ {
		items = append(items, nums[i])
		backtrack(nums, items, ans, i+1)
		items = items[:len(items)-1]
	}
}


