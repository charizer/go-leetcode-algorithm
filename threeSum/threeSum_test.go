package threeSum

import "testing"

func TestThreeSum(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	wants := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	gots := ThreeSum(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}

	arr = []int{0,0,0}
	wants = [][]int{{0, 0, 0}}
	gots = ThreeSum(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}
}

func TestThreeSum2(t *testing.T) {
	arr := []int{-2, 0, 0, 2, 2}
	wants := [][]int{{-2, 0, 2}}
	gots := ThreeSum2(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}
	/*arr := []int{-1, 0, 1, 2, -1, -4}
	wants := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	gots := ThreeSum2(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}*/

	/*arr = []int{0,0,0}
	wants = [][]int{{0, 0, 0}}
	gots = ThreeSum2(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}

	arr = []int{1,1,-2}
	wants = [][]int{{-2, 1, 1}}
	gots = ThreeSum2(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}*/
}

func TestThreeSum3(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	wants := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	gots := ThreeSum3(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}

	arr = []int{0,0,0}
	wants = [][]int{{0, 0, 0}}
	gots = ThreeSum3(arr)
	for idx, got := range gots {
		for i, g := range got {
			if g != wants[idx][i] {
				t.Errorf("idx:%d i:%d got:%d want:%d", idx, i, g, wants[idx][i])
			}
		}
	}
}
