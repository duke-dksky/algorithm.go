package main

import "testing"

func TestTest(t *testing.T) {
	tests := []struct {
		give int
		want int
	}{
		{
			give: 0,
			want: 0,
		},
		{
			give: 1,
			want: 1,
		},
		{
			give: 10,
			want: 2,
		},
		{
			give: 0x7FFFFFFF,
			want: 31,
		},
		{
			give: 0xFFFFFFFF,
			want: 32,
		},
		{
			give: 0x80000000,
			want: 1,
		},
	}
	for _, tt := range tests {
		actual := NumberOf1_Solution1(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
