package arrayslices

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result %d, expect %d", result, expect)
	}
}
