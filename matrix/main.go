package main

import "fmt"

func gen(width, height int) [][]int {
	result := make([][]int, height)
	for k := range result {
		result[k] = make([]int, width)
		if k == 0 {
			result[k][width-1] = 1
			continue
		}
		for i := width - 1; i >= 0; i-- {
			var sum int
			for l := k - 1; l >= 0; l-- {
				sum += result[l][i]
			}
			for l := i; l < width; l++ {
				sum += result[k][l]
			}
			result[k][i] = sum
		}
	}
	return result
}

func output(a [][]int) {
	for _, row := range a {
		for _, v := range row {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
}

func main() {
	var width, height int
	fmt.Print("Enter width and height: ")
	fmt.Scan(&width, &height)
	fmt.Println("Result:")
	output(gen(width, height))
}
