package move_zeroes

func MoveZeroes(nums []int)  {
	//下一个不为0的存放位置
	left := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}
