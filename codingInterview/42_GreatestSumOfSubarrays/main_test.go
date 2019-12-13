package main

import "testing"

func TestFindGreatestSumOfSubArray(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{1, -2, 3, 10, -4, 7, 2, -5},
			want: 18,
		},
		{
			give: []int{-2, -8, -1, -5, -9},
			want: -1,
		},
		{
			give: []int{2, 8, 1, 5, 9},
			want: 25,
		},
	}
	for _, tt := range tests {
		actual := FindGreatestSumOfSubArray(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
