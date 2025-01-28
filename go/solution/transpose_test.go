package solution

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{
			name: "5x5 matrix transpose",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			output: [][]int{
				{1, 6, 11, 16, 21},
				{2, 7, 12, 17, 22},
				{3, 8, 13, 18, 23},
				{4, 9, 14, 19, 24},
				{5, 10, 15, 20, 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transpose(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("Transpose() = %v, want %v", got, tt.output)
			}
		})
	}
}
