package main

import "testing"

func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{0, 3}, []int{}},
		{[]int{3, 10}, []int{10}},
		{[]int{13, 4}, []int{13, 4}},
		{[]int{-3, -11}, []int{-11}},
		{[]int{-3, 2}, []int{2}},
		{[]int{-10, -5}, []int{-10, -5}},
		{[]int{-11, 7}, []int{-11, 7}},
	} {
		got := Filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}

	}
}

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
