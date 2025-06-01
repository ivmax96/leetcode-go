// problem 69 https://leetcode.com/problems/sqrtx/
package easy

func mySqrt(x int) int {
	for i := 0; ; i++ {
		p := i * i
		if p == x {
			return i
		}
		if p > x {
			return i - 1
		}
	}
}
