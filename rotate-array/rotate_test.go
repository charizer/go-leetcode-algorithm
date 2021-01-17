package rotate_array

import "testing"

func TestRotate(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	want := []int{5,6,7,1,2,3,4}
	Rotate(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{-1,-100,3,99}
	k = 2
	want = []int{3,99,-1,-100}
	Rotate(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{1,2,3,4,5,6,7}
	k = 3
	want = []int{5,6,7,1,2,3,4}
	Rotate2(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{-1,-100,3,99}
	k = 2
	want = []int{3,99,-1,-100}
	Rotate2(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
}

func TestRotate3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	want := []int{5, 6, 7, 1, 2, 3, 4}
	Rotate3(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{-1, -100, 3, 99}
	k = 2
	want = []int{3, 99, -1, -100}
	Rotate3(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
}

func TestRotate4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	want := []int{5, 6, 7, 1, 2, 3, 4}
	Rotate4(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{-1, -100, 3, 99}
	k = 2
	want = []int{3, 99, -1, -100}
	Rotate4(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
}
