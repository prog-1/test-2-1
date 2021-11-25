package main

import "fmt"

func filter(s []int) []int {
	for i := range s {
		if s[i]%3 == 0 {
			s = s[:len(s)-1]
		}
	}
	return s
}

func main() {
	fmt.Println("Enter the number of elements in a slice:")
	var size uint
	fmt.Scan(&size)
	s := make([]int, size)
	fmt.Println("Enter the elements:")
	for i := range s {
		fmt.Scan(&s[i])
	}
	an := filter(s)
	fmt.Printf("The filtered slice: %v", an)
}
