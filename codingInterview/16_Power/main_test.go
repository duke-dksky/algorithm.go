package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		give_1 int
		give_2 int
		want   float64
	}{
		{
			give_1: 2,
			give_2: 3,
			want:   8,
		},
		{
			give_1: -2,
			give_2: 3,
			want:   -8,
		},
		{
			give_1: 2,
			give_2: 0,
			want:   1,
		},
		{
			give_1: 0,
			give_2: 0,
			want:   1,
		},
		{
			give_1: 0,
			give_2: 4,
			want:   0,
		},
		{
			give_1: 0,
			give_2: -4,
			want:   0,
		},
	}
	for _, tt := range tests {
		actual := Power(tt.give_1, tt.give_2)
		if actual != tt.want {
			t.Errorf("Add give: %v.%v, actual: %f, want: %f", tt.give_1, tt.give_2, actual, tt.want)
		}
	}

}
