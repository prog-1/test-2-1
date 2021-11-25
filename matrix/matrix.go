package main

import "fmt"

func gen(width, height int) [][]int {
	s := make([][]int, height)
	for i := range s {
		s[i] = make([]int, width)
	}
	s[0][width-1] = 1
	for r := 1; r < height; r++ {
		for c := width - 1; c >= 0; c-- {
			for c1 := c + 1; c1 < width; c1++ {
				s[r][c] += s[r][c1]
			}
			for r1 := 0; r1 < r; r1++ {
				s[r][c] += s[r1][c]
			}
		}
	}
	return s
}

func main() {
	var width, height int
	fmt.Println("Enter width and heigth:")
	fmt.Scan(&width, &height)
	a := gen(width, height)
	for _, row := range a {
		for _, v := range row {
			fmt.Printf("%5d", v)
		}
		fmt.Println()
	}
}
