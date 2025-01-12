package solution

import (
	"slices"
)

type Pair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}

	for _, v := range nums {
		frequency[v]++
	}

	pl := []Pair{}

	for k, v := range frequency {
		pl = append(pl, Pair{k, v})
	}

	slices.SortFunc(pl, func(a, b Pair) int {
		return b.Value - a.Value
	})

	rank := make([]int, k)
	for i := range k {
		rank[i] = pl[i].Key
	}
	return rank
}
