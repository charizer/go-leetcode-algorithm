package permutations

func Permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 0 {
		return ans
	}
	path := []int{}
	backtrack(nums, path, &ans)
	return ans
}

func contain(path []int, target int) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == target {
			return true
		}
	}
	return false
}

func backtrack(nums, path []int, ans *[][]int){
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if contain(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		backtrack(nums, path, ans)
		path = path[0:len(path)-1]
 	}
}


func Permute2(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	backtrack2(&ans, nums, path)
	return ans
}

func backtrack2(ans *[][]int, nums, path []int){
	if len(path) == len(nums) {
		*ans = append(*ans, path)
		return
	}
	for i:= 0; i < len(nums); i++ {
		if contain(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		backtrack2(ans, nums, path)
		path = path[0:len(path)-1]
	}
}


func Permute3(nums []int) [][]int {
	ans := [][]int{}
	backtrack3(0, len(nums)-1, nums, &ans)
	return ans
}

func backtrack3(p, q int, nums []int, ans *[][]int) {
	if p == q {
		*ans = append(*ans, append([]int{}, nums...))
		return
	}
	for i := p; i <= q; i++ {
		nums[p], nums[i] = nums[i], nums[p]
		backtrack3(p+1, q, nums, ans)
		nums[p], nums[i] = nums[i], nums[p]
	}
}
