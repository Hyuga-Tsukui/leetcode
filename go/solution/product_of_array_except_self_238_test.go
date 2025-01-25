package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			t.Parallel()

			result := productExceptSelf(tt.input)

			if !reflect.DeepEqual(result, tt.output) {
				t.Fatalf("expected %v. got=%v", tt.output, result)
			}
		})
	}
}
