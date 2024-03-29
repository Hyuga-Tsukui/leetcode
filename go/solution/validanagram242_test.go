package solution

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	t.Parallel()
	params := []struct {
		input1 string
		input2 string
		got    bool
	}{
		{"anagram", "nagaram", true}, {"car", "rat", false}, {"aa", "a", false}, {"aac", "acc", false},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			actual := isAnagram(p.input1, p.input2)
			if actual != p.got {
				t.Logf("want %v but got %v", p.got, actual)
				t.Fail()
			}
		})
	}
}
