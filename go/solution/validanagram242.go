package solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if c, ok := m[r]; ok && c > 0 {
			m[r]--
			continue
		} else {
			return false
		}
	}
	return true
}
