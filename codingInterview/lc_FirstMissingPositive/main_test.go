package main

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{1, 2, 0},
			want: 3,
		},
		{
			give: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			give: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}
	for _, tt := range tests {
		actual := FirstMissingPositive(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
