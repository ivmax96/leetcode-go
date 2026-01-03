package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a  b c", 1},
		{"longword", 8},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := lengthOfLastWord(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLastWord(%s) = %d, but want %d\n", tt.s, got, tt.want)
			}
		})
	}
}
