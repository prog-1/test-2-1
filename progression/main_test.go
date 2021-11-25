package main

import "testing"

func TestNthTerm(t *testing.T) {
	for _, tc := range []struct {
		a    uint
		want int
	}{
		{1, 1},
		{0, 0},
		{4, 20},
		{2, 2},
		{3, 7},
	} {
		got := nthTerm(tc.a)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.a, got, tc.want)
		}
	}
}
