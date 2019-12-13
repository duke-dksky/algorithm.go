package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		giveN1 int
		giveN2 int
		want   int
	}{
		{
			giveN1: 1,
			giveN2: 2,
			want:   3,
		},
		{
			giveN1: -2,
			giveN2: -8,
			want:   -10,
		},
		{
			giveN1: -1,
			giveN2: 2,
			want:   1,
		},
	}
	for _, tt := range tests {
		actual := Add(tt.giveN1, tt.giveN2)
		if actual != tt.want {
			t.Errorf("Add giveN1: %v, giveN2: %v, actual: %d, want: %d", tt.giveN1, tt.giveN2, actual, tt.want)
		}
	}

}
