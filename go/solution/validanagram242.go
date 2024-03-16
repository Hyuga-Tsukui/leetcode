package solution

import (
	"maps"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := map[byte]int{}
	b := map[byte]int{}
	for i := 0; i < len(s); i++ {
		a[s[i]]++
		b[t[i]]++
	}
	return maps.Equal(a, b)
}
