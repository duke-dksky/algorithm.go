package main

import "testing"

func TestIsPopOrder(t *testing.T) {
	tests := []struct {
		give_push []int
		give_pop  []int
		want      bool
	}{
		{
			give_push: []int{1, 2, 3, 4, 5},
			give_pop:  []int{4, 3, 5, 1, 2},
			want:      false,
		},
		{
			give_push: []int{1, 2, 3, 4, 5},
			give_pop:  []int{3, 5, 4, 2, 1},
			want:      false,
		},
	}
	for _, tt := range tests {
		actual := IsPopOrder(tt.give_push, tt.give_pop)
		if actual != tt.want {
			t.Errorf("give_push: %v, give_pop: %v, actual: %d, want: %d", tt.give_push, tt.give_pop, actual, tt.want)
		}
	}

}
