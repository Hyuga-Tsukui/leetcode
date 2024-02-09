package solution

import (
	"fmt"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	t.Parallel()
	params := []struct {
		input []int
		got   bool
	}{
		{[]int{1, 2, 3, 1}, true}, {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true}, {[]int{1, 2, 3, 4}, false},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			actual := containsDuplicate(p.input)
			if actual != p.got {
				t.Logf("want %v but got %v", p.got, actual)
				t.Fail()
			}
		})
	}
}
