package main

import (
	"testing"
)

func TestMinNumberInRotatedArray(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			[]int{3, 4, 5, 1, 1, 2},
			1,
		},
		{
			[]int{3, 4, 5, 1, 2, 2},
			1,
		},
		{
			[]int{1, 0, 1, 1, 1},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
		},
		{
			[]int{2},
			2,
		},
	}
	for _, tt := range tests {
		actual := MinNumberInRotatedArray(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
