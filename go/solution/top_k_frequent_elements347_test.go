package solution

import (
	"fmt"
	"slices"
	"testing"
)

func Test_topKFrequnet(t *testing.T) {
	t.Parallel()
	params := []struct {
		input1 []int
		input2 int
		got    []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2, []int{2, -1}},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			actual := topKFrequent(p.input1, p.input2)
			slices.Sort(actual)
			slices.Sort(p.got)

			if !slices.Equal(actual, p.got) {
				t.Logf("want %v but got %v", p.got, actual)
				t.Fail()
			}
		})
	}
}
