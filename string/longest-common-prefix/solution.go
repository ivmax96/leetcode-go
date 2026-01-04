/*
link: https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strings.Builder{}
	for i, ch := range strs[0] {
		for _, s := range strs[1:] {
			if len(s) == i || rune(s[i]) != ch {
				return prefix.String()
			}
		}
		prefix.WriteRune(ch)
	}
	return prefix.String()
}
