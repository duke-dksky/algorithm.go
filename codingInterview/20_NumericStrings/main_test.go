package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{
			give: "100",
			want: true,
		},
		{
			give: "-1E-16",
			want: true,
		},
		{
			give: "3.1416",
			want: true,
		},
		{
			give: "12e+5.4",
			want: false,
		},
	}
	for _, tt := range tests {
		actual := isNumeric(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %v, want: %v", tt.give, actual, tt.want)
		}
	}

}
