// https://leetcode.com/problems/powx-n/

package binary_search

func AbsInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func isNegative(n int) bool{
    if n<0{
        return true
    }
    return false
}
func myPow(x float64, n int) float64 {
    neg := isNegative(n)
    n = AbsInt(n)
    ans := 1.0
    mul := x
    for n > 0 {
        if n%2==1{
            ans = ans * mul;
        }
        if n>1{
        mul = mul*mul
        }
        n = n/2
    }
    if neg {
        return 1.0/ans;
    }
    return ans;
}