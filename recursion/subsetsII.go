// https://leetcode.com/problems/subsets-ii
/* Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package recursion

import (
	"sort"
)

func helperSubsetWithDups(idx int, nums []int, curr []int, ans *[][]int) {
	temp := append([]int(nil), curr...)
	*ans = append(*ans, temp)

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		curr = append(curr, nums[i])
		helperSubsetWithDups(i+1, nums, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var curr []int
	sort.Ints(nums)
	helperSubsetWithDups(0, nums, curr, &ans)
	return ans
}
