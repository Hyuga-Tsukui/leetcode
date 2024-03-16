package solution

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

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
