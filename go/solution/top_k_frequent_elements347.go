package solution

func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}

	for _, v := range nums {
		frequency[v]++
	}
	bucket := make([][]int, len(nums)+1)
	for k, v := range frequency {
		bucket[v] = append(bucket[v], k)
	}
	rank := []int{}
	for i := len(bucket) - 1; i > 0 && len(rank) < k; i-- {
		if len(bucket[i]) > 0 {
			rank = append(rank, bucket[i]...)
		}
	}
	return rank
}
