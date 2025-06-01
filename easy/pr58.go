// Problem: https://leetcode.com/problems/length-of-last-word/

package easy

func lengthOfLastWord(s string) int {
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && res == 0 {
			continue
		}
		if s[i] == ' ' {
			break
		}
		res++
	}
	return res
}
