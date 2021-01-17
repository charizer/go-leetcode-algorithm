package removeDuplicates

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	//记录前一个位置
	pre := 0
	for cur := 1; cur < len(nums); {
		//当前和前一个位置数据重复，则将当前元素删除，有元素删除，不更新pre、cur位置
		if nums[cur] == nums[pre] {
			nums = append(nums[0:cur], nums[cur+1:]...)
		}else{
			//不重复，则更新pre、cur位置
			pre = cur
			cur++
		}
	}
	return len(nums)
}

func RemoveDuplicates3(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 双指针处理
	// 左指针
	left := 0
	for right := 1; right < len(nums); right++ {
		// 相等，则不处理，右指针继续右移
		if nums[right] == nums[left] {
			continue
		}else{
			// 不相等, 则将元素放入left+1的位置，并更新left
			if right - left > 1 {
				nums[left+1] = nums[right]
			}
			left++
		}
	}
	return left + 1
}

func RemoveDuplicates4(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 双指针解法
	left := 0
	for right := 1 ; right < len(nums); right++ {
		// 如果相同的话，继续移动右指针
		if nums[right] == nums[left] {
			continue
		}
		//否则的话，把right的值填入left的下一个位置，并且更新左指针
		left++
		nums[left] = nums[right]
	}
	return left + 1
}
