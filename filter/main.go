package main

import "fmt"

func Filter(s []int) (d []int) {
	for _, v := range s {
		if v%3 != 0 {
			d = append(d, v)
		}
	}
	return d
}
func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(Filter(input))

}
