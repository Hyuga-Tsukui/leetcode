package solution

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	groups := map[[26]int][]string{}

	for _, str := range strs {
		alpha := [26]int{}
		for _, r := range str {
			alpha[r-'a']++
		}
		if group, ok := groups[alpha]; ok {
			groups[alpha] = append(group, str)
		} else {
			groups[alpha] = []string{str}
		}
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
