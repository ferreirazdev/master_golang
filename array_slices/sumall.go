package arrayslices

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, numbers := range nums {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
