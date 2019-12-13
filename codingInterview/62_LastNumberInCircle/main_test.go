package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		give_1 int
		give_2 int
		want   int
	}{
		{
			give_1: 5,
			give_2: 3,
			want:   3,
		},
	}
	for _, tt := range tests {
		actual := LastRemaining(tt.give_1, tt.give_2)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give_1, actual, tt.want)
		}
	}

}
