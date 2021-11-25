package main

import "fmt"

func nthTerm(n uint) int{
	a1 := n-2
	a2 := n-1
	n = 3 * a1 + 2 * a2
	return int(n)
} 

func main(){
	var n uint
	fmt.Println("Enter n number")
	fmt.Scan(&n)
	fmt.Println("Answer is",nthTerm(n))
}