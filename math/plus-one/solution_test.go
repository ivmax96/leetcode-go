package main

import (
	"slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			digits: []int{4, 3, 9, 9},
			want:   []int{4, 4, 0, 0},
		},
		{
			digits: []int{9, 9, 9, 9},
			want:   []int{1, 0, 0, 0, 0},
		},
		{
			digits: []int{9},
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := plusOne(tt.digits)
			if !slices.Equal(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, but want %v\n", tt.digits, got, tt.want)
			}
		})
	}
}
