package main

import "testing"

func Testfilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{3}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4}},
		{[]int{2, 1, 8}, []int{2, 1, 8}},
		{[]int{3, 7, 9}, []int{7}},
		{[]int{-3, 1, 8}, []int{1, 8}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		filter(got)
		if !equal(got, tc.want) {
			t.Errorf("filter(%v) = %v, want = %v", tc.s, got, tc.want)
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
