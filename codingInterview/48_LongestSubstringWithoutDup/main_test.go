package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{
			give: "acfrarabc",
			want: 4,
		},
		{
			give: "acfrarabc",
			want: 4,
		},
		{
			give: "arabcacfr",
			want: 4,
		},
		{
			give: "aaaa",
			want: 1,
		},
		{
			give: "abcdefg",
			want: 7,
		},
		{
			give: "aaabbbccc",
			want: 2,
		},
		{
			give: "abcdcba",
			want: 4,
		},
		{
			give: "abcdaef",
			want: 6,
		},
		{
			give: "a",
			want: 1,
		},
		{
			give: "",
			want: 0,
		},
	}
	for _, tt := range tests {
		actual := longestSubstringWithoutDuplication(tt.give)
		if actual != tt.want {
			t.Errorf("Add give: %v, actual: %d, want: %d", tt.give, actual, tt.want)
		}
	}

}
