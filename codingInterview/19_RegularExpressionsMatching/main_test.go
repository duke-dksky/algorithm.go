package main

import "testing"

func TestMatch(t *testing.T) {
	tests := []struct {
		give_1 string
		give_2 string
		want   bool
	}{
		{
			give_1: "aaa",
			give_2: "aaa*",
			want:   true,
		},
		{
			give_1: "aaa",
			give_2: "ab*a*c*a",
			want:   true,
		},
		{
			give_1: "aaa",
			give_2: "ab*ac*a",
			want:   true,
		},
	}
	for _, tt := range tests {
		actual := Match(tt.give_1, tt.give_2)
		if actual != tt.want {
			t.Errorf("Add give: %v, %v, actual: %v, want: %v", tt.give_1, tt.give_2, actual, tt.want)
		}
	}

}
