package arrayslices

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}

	result := Sum(nums)
	expect := 15

	if result != expect {
		t.Errorf("result %d, expect %d, dado %v", result, expect, nums)
	}
}
