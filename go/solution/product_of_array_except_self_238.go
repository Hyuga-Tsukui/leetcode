package solution

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// initialize.
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(result); j++ {
			if j == i {
				continue
			}
			result[j] = result[j] * nums[i]
		}
	}

	return result
}
