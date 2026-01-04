package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"hello", "helloworld"}, "hello"},
		{[]string{""}, ""},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("lengthOfLastWord(%v) = %q, but want %q\n", tt.strs, got, tt.want)
			}
		})
	}
}
