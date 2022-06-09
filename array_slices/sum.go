package arrayslices

func Sum(nums [5]int) int {
	soma := 0
	for _, num := range nums {
		soma += num
	}

	return soma
}
