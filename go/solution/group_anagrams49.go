package solution

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	gropus := map[string][]string{}

	for _, v := range strs {
		splited := strings.Split(v, "")
		slices.Sort(splited)
		joined := strings.Join(splited, "")
		if s, ok := gropus[joined]; !ok {
			gropus[joined] = []string{v}
		} else {
			gropus[joined] = append(s, v)
		}
	}

	result := [][]string{}

	for _, v := range gropus {
		result = append(result, v)
	}
	return result
}
