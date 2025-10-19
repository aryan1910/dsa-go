// https://leetcode.com/problems/accounts-merge
package graph

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	ds := NewDisjointSet(len(accounts))
	mapEmailToID := make(map[string]int)
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			email := accounts[i][j]
			if _, ok := mapEmailToID[email]; !ok {
				mapEmailToID[email] = i
			} else {
				ds.UnionByRank(i, mapEmailToID[email])
				mapEmailToID[email] = i
			}
		}
	}
	names := make(map[int][]string)
	for key, val := range mapEmailToID {
		ulp := ds.FindPar(val)
		names[ulp] = append(names[ulp], key)
	}
	var result [][]string
	for key, val := range names {
		sort.Strings(val)
		curr := []string{accounts[key][0]}
		curr = append(curr, val...)
		result = append(result, curr)
	}
	return result

}
