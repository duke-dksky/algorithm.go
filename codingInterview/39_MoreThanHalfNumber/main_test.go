package main

import "testing"

func TestMoreThanHalfNum(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			want: 2,
		},
		{
			give: []int{2, 2, 2, 2, 2, 1, 3, 4, 5},
			want: 2,
		},
		{
			give: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			want: 2,
		},
		{
			give: []int{1, 3, 4, 5, 2, 2, 2, 2, 2},
			want: 2,
		},
		{
			give: []int{1, 2, 3, 2, 4, 2, 5, 2, 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		actual := MoreThanHalfNum(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
