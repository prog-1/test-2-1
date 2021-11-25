package main

import (
	"reflect"
	"testing"
)

func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
func TestGen(t *testing.T) {
	for _, tc := range []struct {
		width, height int
		want          [][]int
	}{
		{1, 1, [][]int{
			{1},
		}},
		{4, 5, [][]int{
			{0, 0, 0, 1},
			{4, 2, 1, 1},
			{16, 7, 3, 2},
			{53, 21, 8, 4},
			{159, 58, 20, 8},
		}},
		{2, 10, [][]int{
			{0, 1},
			{1, 1},
			{3, 2},
			{8, 4},
			{20, 8},
			{48, 16},
			{112, 32},
			{256, 64},
			{576, 128},
			{1280, 256},
		}},
	} {
		got := Gen(tc.width, tc.height)
		if !equal(got, tc.want) {
			t.Errorf("gen(%v, %v) = %v, want = %v", tc.width, tc.height, got, tc.want)
		}
	}
}
