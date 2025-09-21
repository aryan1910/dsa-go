// https://leetcode.com/problems/permutation-sequence
/*
The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
Input: n = 3, k = 3
Output: "213"
*/

package recursion

import (
	"fmt"
	"strings"
)

func getPermutation(n int, k int) string {
	fact := 1
	nums := make([]int, n)
	for i := 1; i < n; i++ {
		fact = fact * i
		nums[i-1] = i
	}
	k = k - 1
	nums[n-1] = n
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(fmt.Sprintf("%d", nums[k/fact]))
		nums = append(nums[:(k/fact)], nums[(k/fact)+1:]...)
		if len(nums) == 0 {
			break
		}
		k = k % fact
		fact = fact / len(nums)
	}
	return sb.String()
}
