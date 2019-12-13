package main

import "testing"

func TestGetMaxValue(t *testing.T) {
	tests := []struct {
		give [][]int
		want int
	}{
		{
			give: [][]int{
				{1, 10, 3, 8},
				{12, 2, 9, 6},
				{5, 7, 4, 11},
				{3, 7, 16, 5},
			},
			want: 53,
		},
		{
			give: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		actual := GetMaxValue_solution3(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
