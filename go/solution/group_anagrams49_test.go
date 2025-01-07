package solution

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"testing"
)

func Test_group_anagrams49(t *testing.T) {
	t.Parallel()
	params := []struct {
		input []string
		got   [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{"", ""}, [][]string{{"", ""}}},
	}
	for _, p := range params {
		t.Run(fmt.Sprintf("Testing [%v]", p), func(t *testing.T) {
			actual := groupAnagrams(p.input)
			if !HelperCompareAnagramGroups(actual, p.got) {
				t.Logf("want %v but got %v", p.got, actual)
				t.Fail()
			}
		})
	}
}

func HelperCompareAnagramGroups(a, b [][]string) bool {
	sortFn := func(grpuos [][]string) {
		for _, g := range grpuos {
			slices.Sort(g)
		}
		sort.Slice(grpuos, func(i, j int) bool {
			return grpuos[i][0] < grpuos[j][0]
		})
	}
	sortFn(a)
	sortFn(b)
	return reflect.DeepEqual(a, b)
}
