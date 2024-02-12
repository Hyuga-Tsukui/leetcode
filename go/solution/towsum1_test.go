package solution

import (
	"fmt"
	"slices"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Parallel()
	params := []struct {
		input1 []int
		input2 int
		got    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}}, {[]int{3, 2, 4}, 6, []int{1, 2}}, {[]int{3, 3}, 6, []int{0, 1}}, {[]int{-1, -3, 4}, 3, []int{0, 2}},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			actual := twoSum(p.input1, p.input2)
			if !slices.Equal(actual, p.got) {
				t.Logf("want %v but got %v", p.got, actual)
				t.Fail()
			}
		})
	}
}
