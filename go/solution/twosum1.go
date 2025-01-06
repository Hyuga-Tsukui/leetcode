package solution

func twoSum(nums []int, target int) []int {
	if len(nums) <= 2 {
		return []int{0, 1}
	}
	m := map[int]int{}
	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{v, i}
		}
		m[n] = i
	}
	return []int{}
}
