package main

import "fmt"

func nthTerm(n uint) int {
	var a, b int
	for n != n-2 {
		for n != n-1 {
			b++
			b = b * 3
		}
		a++
		a = a * 3
	}
	a = a + b
	return a
}

func main() {
	fmt.Println("Enter n:")
	var n uint
	fmt.Scan(n)
	an := uint(nthTerm(n))
	fmt.Printf("The nth term is %v", an)
}
