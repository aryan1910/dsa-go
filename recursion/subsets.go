// https://leetcode.com/problems/subsets/
/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package recursion

import (
	"sort"
)

func helperSubsets(idx int, nums []int, curr []int, ans *[][]int) {
	if idx == len(nums) {
		*ans = append(*ans, append([]int(nil), curr...))
		return
	}
	// Skip
	helperSubsets(idx+1, nums, curr, ans)

	// Keep
	curr = append(curr, nums[idx])
	helperSubsets(idx+1, nums, curr, ans)
	curr = curr[:len(curr)-1]
}

func helperSubsetsUsingFor(idx int, nums []int, curr []int, ans *[][]int) {
	*ans = append(*ans, append([]int(nil), curr...))
	for i := idx; i < len(nums); i++ {
		// This can be skipped for this problem since unique elements are ensured in input
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		curr = append(curr, nums[i])
		helperSubsets(i+1, nums, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

func subsets(nums []int) [][]int {
	var ans [][]int
	var curr []int
	sort.Ints(nums)
	helperSubsets(0, nums, curr, &ans)
	return ans
}
