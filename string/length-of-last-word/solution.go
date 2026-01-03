/*
link: https://leetcode.com/problems/length-of-last-word/description/

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/
package main

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	lastWord := words[len(words)-1]
	return len(lastWord)
}
