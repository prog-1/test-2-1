package main

import (
	"reflect"
	"testing"
)

func Testgen(t *testing.T) {
	for _, tc := range []struct {
		width, height int
		want          [][]int
	}{
		{0, 0, [][]int{}},
		{1, 1, [][]int{
			{1},
		}},
		{5, 7, [][]int{
			{0, 0, 0, 0, 1},
			{8, 4, 2, 1, 1},
			{36, 16, 7, 3, 2},
			{130, 53, 21, 8, 4},
			{419, 159, 58, 20, 8},
			{1257, 448, 152, 48, 16},
			{3586, 1208, 384, 112, 32},
		}},
		{3, 3, [][]int{
			{0, 0, 1},
			{2, 1, 1},
			{7, 3, 2},
		}},
		{4, 2, [][]int{
			{0, 0, 0, 1},
			{4, 2, 1, 1},
		}},
	} {
		got := gen(tc.width, tc.height)
		if !equal(got, tc.want) {
			t.Errorf("gen(%v, %v) = %v, want = %v", tc.width, tc.height, got, tc.want)
		}
	}
}

func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
