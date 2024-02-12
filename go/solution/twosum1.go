package solution

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if target-(v1+v2) == 0 {
				return []int{i, j + i + 1}
			}
			continue
		}
	}
	return []int{}
}
