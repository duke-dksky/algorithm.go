package main

import (
	"testing"
)

func TestGetNumberOfK(t *testing.T) {
	tests := []struct {
		giveData []int
		giveK    int
		want     int
	}{
		{
			giveData: []int{1, 2, 3, 3, 3, 3, 4, 5},
			giveK:    3,
			want:     4,
		},
		{
			giveData: []int{1, 2, 4, 5},
			giveK:    6,
			want:     0,
		},
		{
			giveData: []int{1},
			giveK:    1,
			want:     1,
		},
		{
			giveData: []int{},
			giveK:    1,
			want:     0,
		},
	}

	for _, tt := range tests {
		actual := GetNumberOfK(tt.giveData, tt.giveK)
		if actual != tt.want {
			t.Errorf("actual: %d, expected: %d", actual, tt.want)
		}
	}
}
