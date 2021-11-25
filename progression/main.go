package main

import "fmt"

func nthTerm(n uint) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return 3*nthTerm(n-2) + 2*nthTerm(n-1)
}

func main() {
	fmt.Print("Enter n: ")
	var s uint
	fmt.Scan(&s)
	fmt.Print("The nth term is ", nthTerm(s), ".")
}
