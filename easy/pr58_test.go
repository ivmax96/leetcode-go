package easy

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"somestring", 10},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := lengthOfLastWord(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLastWord(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
