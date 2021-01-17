package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	// 从尾部插入
	// 用于记录数据放入的位置
	p1, p2, tail := m-1, n-1, m+n-1
	// nums2没处理完就需要继续处理
	for p2 >= 0 {
		// nums1和nums2都没有处理完，从大到小插入nums1和nums2元素都nums1结尾
		if p1 >= 0 {
			if nums1[p1] > nums2[p2] {
				nums1[tail] = nums1[p1]
				p1--
			}else{
				nums1[tail] = nums2[p2]
				p2--
			}
		// nums1已经处理完，再把剩余的nums2插入到nums1中
		}else{
			nums1[tail] = nums2[p2]
			p2--
		}
		tail--
	}
}
