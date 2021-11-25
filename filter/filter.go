package main

import "fmt"

func filter(s []int) (r []int) {
	for _, v := range s {
		if v%3 != 0 {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	fmt.Println("Enter the slice size:")
	var size int
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println("The filtered slice:", filter(input))
}
