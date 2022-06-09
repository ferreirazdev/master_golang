package arrayslices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection with size of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		result := Sum(nums)
		expect := 15

		if result != expect {
			t.Errorf("result %d, expect %d, dado %v", result, expect, nums)
		}
	})

	t.Run("collection with size of any numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}

		result := Sum(nums)
		expect := 6

		if result != expect {
			t.Errorf("result %d, expect %d, dado %v", result, expect, nums)
		}
	})

}
