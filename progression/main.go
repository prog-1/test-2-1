package main

import "fmt"

func NthTerm(n uint) int {
	a, b := 0, 1
	var i uint
	for ; i < n; i++ {
		a, b = b, 3*a+2*b

	}
	return a
}
func main() {
	var n uint
	fmt.Scan(&n)
	fmt.Println(NthTerm(n))
}
