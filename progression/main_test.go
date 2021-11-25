package main

import "testing"

func TestnthTerm(t *testing.T) {
	for _, tc := range []struct {
		n    uint
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 7},
		{4, 20},
	} {
		got := nthTerm(tc.n)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
