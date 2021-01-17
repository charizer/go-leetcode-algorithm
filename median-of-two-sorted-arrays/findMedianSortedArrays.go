package median_of_two_sorted_arrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		}else{
			nums = append(nums, nums2[j])
			j++
		}
	}
	if i == len(nums1) {
		nums = append(nums, nums2[j:]...)
	}
	if j == len(nums2) {
		nums = append(nums, nums1[i:]...)
	}
	if len(nums) % 2 == 1 {
		return float64(nums[(len(nums)-1)/ 2])
	}else{
		return float64(nums[(len(nums)-1) / 2 ] + nums[len(nums) / 2]) / 2
	}
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	tn := m + n
	ap, bp, left, right := 0, 0, 0, 0
	for i := 0; i <= tn / 2 ; i++ {
		left = right
		if ap < m && (bp > n || nums1[ap] < nums2[bp]) {
			right = nums1[ap]
			ap++
		}else{
			right = nums2[bp]
			bp++
		}
	}
	if (tn & 1) == 0 {
		return float64(left + right) / 2.0
	} else {
		return float64(right)
	}
}


func FindMedianSortedArrays3(nums1 []int, nums2 []int) float64 {

}
