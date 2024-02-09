package solution

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	params := []struct {
		input1 string
		input2 string
		got    bool
	}{
		{"anagram", "nagaram", true}, {"car", "rat", false}, {"aa", "a", false}, {"aac", "acc", false},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			acutual := isAnagram(p.input1, p.input2)
			if acutual != p.got {
				t.Logf("want %v but got %v", p.got, acutual)
				t.Fail()
			}
		})
	}
}
