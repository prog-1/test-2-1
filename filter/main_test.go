package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{3, 6, 8}, []int{8}},
		{[]int{-3, -6, -8}, []int{-8}},
		{[]int{6, 0, 9}, []int{}},
		{[]int{1, 4, 2}, []int{1, 4, 2}},
		{[]int{2, 6, 8}, []int{2, 8}},
	} {
		got := filter(tc.s)
		if !Equal(got, tc.want) {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
