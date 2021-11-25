package main

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
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
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{0, 3}, []int{}},
		{[]int{3, 2}, []int{2}},
		{[]int{14, 56}, []int{14, 56}},
		{[]int{3, 9, 18}, []int{}},
		{[]int{-3, -7}, []int{-7}},
		{[]int{-3, 10}, []int{10}},
		{[]int{-360, -3}, []int{}},
	} {
		got := Filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}

	}
}
