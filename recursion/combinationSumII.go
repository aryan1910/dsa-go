// https://leetcode.com/problems/combination-sum-ii/
/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
*/
package recursion

import "sort"

func helperCombinationSum2(idx int, c []int, target int, curr []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int(nil), curr...))
		return
	}
	if target < 0 || idx == len(c) {
		return
	}
	for i := idx; i < len(c); i++ {
		if i > idx && c[i] == c[i-1] {
			continue
		}
		if target < 0 {
			break
		}
		curr = append(curr, c[i])
		helperCombinationSum2(i+1, c, target-c[i], curr, ans)
		curr = curr[:len(curr)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int
	var curr []int

	helperCombinationSum2(0, candidates, target, curr, &ans)
	return ans
}
