package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		give int
		want int
	}{
		{
			give: 5,
			want: 15,
		},
	}
	for _, tt := range tests {
		actual := Sum(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
