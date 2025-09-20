// https://leetcode.com/problems/combination-sum/
/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
*/

package recursion

import "sort"

func helperCombinationSum(i int, c []int, target int, curr []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int(nil), curr...))
		return
	}
	if target < 0 || i == len(c) {
		return
	}
	helperCombinationSum(i+1, c, target, curr, ans)

	curr = append(curr, c[i])
	helperCombinationSum(i, c, target-c[i], curr, ans)
	curr = curr[:len(curr)-1]
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int
	var curr []int

	helperCombinationSum(0, candidates, target, curr, &ans)
	return ans
}
