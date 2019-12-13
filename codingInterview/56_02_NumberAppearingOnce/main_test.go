package main

import "testing"

func TestFindNumberAppearingOnce(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{1, 1, 2, 2, 2, 1, 3},
			want: 3,
		},
		{
			give: []int{4, 3, 3, 2, 2, 2, 3},
			want: 4,
		},
		{
			give: []int{4, 4, 1, 1, 1, 7, 4},
			want: 7,
		},
		{
			give: []int{-1024, -1024, -1024, -1023},
			want: -1023,
		},
		{
			give: []int{-23, 0, 214, -23, 214, -23, 214},
			want: 0,
		},
		{
			give: []int{-10, 214, 214, 214},
			want: -10,
		},
	}

	for _, tt := range tests {
		actual := FindNumberAppearingOnce(tt.give)
		if actual != tt.want {
			t.Errorf("FindNumberAppearingOnce give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}
}
