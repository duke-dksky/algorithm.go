package main

import "testing"

func TestNumberOf1Between1AndN(t *testing.T) {
	tests := []struct {
		give int
		want int
	}{
		{
			give: 1,
			want: 1,
		},
		{
			give: 5,
			want: 1,
		},
		{
			give: 10,
			want: 2,
		},
		{
			give: 55,
			want: 16,
		},
		{
			give: 99,
			want: 20,
		},
		{
			give: 10000,
			want: 4001,
		},
		{
			give: 21345,
			want: 18821,
		},
		{
			give: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		actual := NumberOf1Between1AndN(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
