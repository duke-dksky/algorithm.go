package main

import (
	"reflect"
	"testing"
)

func TestFindFirstBitIs1(t *testing.T) {
	tests := []struct {
		give int
		want uint
	}{
		{
			give: 2,
			want: 1,
		},
		{
			give: 4,
			want: 2,
		},
	}

	for _, tt := range tests {
		actual := FindFirstBitIs1(tt.give)
		if actual != tt.want {
			t.Errorf("FindFirstBitIs1 %d actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}

func TestIsBit1(t *testing.T) {
	tests := []struct {
		giveNum      int
		giveIndexBit uint
		want         bool
	}{
		{
			giveNum:      2,
			giveIndexBit: 1,
			want:         true,
		},
	}

	for _, tt := range tests {
		actual := IsBit1(tt.giveNum, tt.giveIndexBit)
		if actual != tt.want {
			t.Errorf("IsBit1 actual: %v, want: %v", actual, tt.want)
		}
	}

}

func TestFindNumsAppearOnce(t *testing.T) {
	tests := []struct {
		give []int
		want []int
	}{
		{
			give: []int{2, 4, 3, 6, 3, 2, 5, 5},
			want: []int{6, 4},
		},
		{
			give: []int{4, 6},
			want: []int{6, 4},
		},
		{
			give: []int{4, 6, 1, 1, 1, 1},
			want: []int{6, 4},
		},
	}
	for _, tt := range tests {
		actual := FindNumsAppearOnce(tt.give)
		if !reflect.DeepEqual(tt.want, actual) {
			t.Errorf("IsBit1 actual: %v, want: %v", actual, tt.want)
		}
	}
}
