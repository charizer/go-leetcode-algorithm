package moveZeroes

func MoveZeroes(nums []int)  {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func MoveZeroes2(nums []int)  {
	left := 0
	for right := 0; right < len(nums); right ++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
