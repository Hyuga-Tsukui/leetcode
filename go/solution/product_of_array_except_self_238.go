package solution

// space complexity O(1)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = 1
	}

	// 先頭の数字の前は存在しないのでSuffixとの計算に備えて1にしておく.
	// この変数直前のPrefixの計算結果を保持する.
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// 末端の数字の後ろは存在しないのでPrefixとの計算に備えて1にしておく.
	// 直前のSuffixの計算結果を保持する
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

// space complexity O(n)
// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	result := make([]int, n)
// 	prefix := make([]int, n)
// 	suffix := make([]int, n)
//
// 	// 先頭のprefixはsuffixとの計算を考えると必ず1になる.
// 	prefix[0] = 1
// 	for i := 1; i < n; i++ {
// 		prefix[i] = nums[i-1] * prefix[i-1]
// 	}
//
// 	// 末端のsuffixは1.
// 	suffix[n-1] = 1
//
// 	for i := n - 2; i >= 0; i-- {
// 		suffix[i] = nums[i+1] * suffix[i+1]
// 	}
//
// 	for i := 0; i < n; i++ {
// 		result[i] = suffix[i] * prefix[i]
// 	}
//
// 	return result
// }

// naive solution
// func productExceptSelf(nums []int) []int {
// 	result := make([]int, len(nums))
//
// 	// initialize.
// 	for i := 0; i < len(result); i++ {
// 		result[i] = 1
// 	}
//
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(result); j++ {
// 			if j == i {
// 				continue
// 			}
// 			result[j] = result[j] * nums[i]
// 		}
// 	}
//
// 	return result
// }
