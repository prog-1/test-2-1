package main

import "testing"

func TestNthTerm(t *testing.T) {
	for _, tc := range []struct {
		n    uint
		want int
	}{

		{n: 4, want: 20},
		{n: 1, want: 1},
		{n: 0, want: 0},
	} {
		got := NthTerm(tc.n)
		if got != tc.want {
			t.Errorf("Sort(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
