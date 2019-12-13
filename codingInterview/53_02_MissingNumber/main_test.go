package main

import (
	"testing"
)

func TestGetMissingNumber(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			give: []int{0},
			want: 1,
		},
		{
			give: []int{1},
			want: 0,
		},
	}

	for _, tt := range tests {
		actual := GetMissingNumber(tt.give)
		if actual != tt.want {
			t.Errorf("GetMissingNumber %v, want: %d, actual: %d.", tt.give, tt.want, actual)
		}
	}
}
