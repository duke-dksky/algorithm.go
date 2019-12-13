package main

import "testing"

func TestGetTranslationCount(t *testing.T) {
	tests := []struct {
		give int
		want int
	}{
		{
			give: 10,
			want: 2,
		},
		{
			give: 125,
			want: 3,
		},
		{
			give: 126,
			want: 2,
		},
		{
			give: 426,
			want: 1,
		},
		{
			give: 100,
			want: 2,
		},
		{
			give: -100,
			want: 0,
		},
	}
	for _, tt := range tests {
		actual := GetTranslationCount(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
