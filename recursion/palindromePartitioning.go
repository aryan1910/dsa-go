// https://leetcode.com/problems/palindrome-partitioning
/*
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.
Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]
*/

package recursion

func isPalindrome(str string, s int, e int) bool {
	for s <= e {
		if str[e] != str[s] {
			return false
		}
		s++
		e--
	}
	return true
}

func helperPartition(idx int, str string, curr []string, ans *[][]string) {
	if idx == len(str) {
		*ans = append(*ans, append([]string(nil), curr...))
		return
	}

	for i := idx; i < len(str); i++ {
		if isPalindrome(str, idx, i) {
			curr = append(curr, string(str[idx:i+1]))
			helperPartition(i+1, str, curr, ans)
			curr = curr[:len(curr)-1]
		}
	}
}

func partition(s string) [][]string {
	var curr []string
	var ans [][]string
	helperPartition(0, s, curr, &ans)
	return ans
}
