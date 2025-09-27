// https://leetcode.com/problems/permutations/
// Given an array nums of distinct integers, return all the possible permutations.
// You can return the answer in any order.

package recursion

func helperPermute(nums []int, freq map[int]bool, curr []int, ans *[][]int) {
	if len(curr) == len(nums) {
		*ans = append(*ans, append([]int(nil), curr...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !freq[i] {
			freq[i] = true
			curr = append(curr, nums[i])
			helperPermute(nums, freq, curr, ans)
			freq[i] = false
			curr = curr[:len(curr)-1]
		}
	}
}

func swapHelperPermute(idx int, nums []int, ans *[][]int) {
	if idx == len(nums) {
		*ans = append(*ans, append([]int(nil), nums...))
		return
	}

	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		swapHelperPermute(idx+1, nums, ans)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}

func permute(nums []int) [][]int {
	// var curr []int
	var ans [][]int
	//freq := make(map[int]bool)

	//helper(nums, freq, curr, &ans)
	swapHelperPermute(0, nums, &ans)
	return ans
}
