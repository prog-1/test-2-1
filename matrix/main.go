package main

import "fmt"

func Gen(height, width int) [][]int {
	height, width = width, height
	rows, cols := height, width
	a := make([][]int, rows)
	for row := 0; row < rows; row++ {
		a[row] = make([]int, cols)
	}

	a[0][width-1] = 1
	for r := 1; r <= height-1; r++ {
		for c := width - 1; c >= 0; c-- {
			for c1 := c + 1; c1 < width; c1++ {
				a[r][c] = a[r][c] + a[r][c1]

			}
			for r1 := 0; r1 < r; r1++ {

				a[r][c] += a[r1][c]
			}

		}

	}
	return a
}
func main() {
	var width, height int
	fmt.Println("Enter width and heigth:")
	fmt.Scan(&width, &height)
	if width == 0 || height == 0 {
		fmt.Println("width and heigth can't be zero or negative")
	}
	table := Gen(width, height)
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%5d", table[i][j])
		}
		fmt.Println()
	}
}
