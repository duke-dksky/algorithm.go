package main

import "testing"

func TestVerifySquenceOfBST(t *testing.T) {
	tests := []struct {
		give []int
		want bool
	}{
		{
			give: []int{4, 8, 6, 12, 16, 14, 10},
			want: true,
		},
		{
			give: []int{4, 6, 7, 5},
			want: true,
		},
		{
			give: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			give: []int{5, 4, 3, 2, 1},
			want: true,
		},
		{
			give: []int{5},
			want: true,
		},
		{
			give: []int{7, 4, 6, 5},
			want: false,
		},
		{
			give: []int{4, 6, 12, 8, 16, 14, 10},
			want: false,
		},
		{
			give: []int{},
			want: false,
		},
	}
	for _, tt := range tests {
		actual := VerifySquenceOfBST(tt.give)
		if actual != tt.want {
			t.Errorf("give: %v, want: %v, actual: %v", tt.give, tt.want, actual)
		}
	}
}
