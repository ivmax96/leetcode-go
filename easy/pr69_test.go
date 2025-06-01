package easy

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{24, 4},
		{1, 1},
		{0, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := mySqrt(tt.x)
			if got != tt.want {
				t.Errorf("mySqrt(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
