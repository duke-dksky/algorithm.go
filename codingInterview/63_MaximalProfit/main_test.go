package main

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{
			give: []int{4, 1, 3, 2, 5},
			want: 4,
		},
		{
			give: []int{1, 2, 4, 7, 11, 16},
			want: 15,
		},
		{
			give: []int{16, 16, 16, 16, 16},
			want: 0,
		},
	}

	for _, tt := range tests {
		actual := MaxDiff(tt.give)
		if actual != tt.want {
			t.Errorf("want: %v, actual: %v.", tt.want, actual)
		}
	}

}
