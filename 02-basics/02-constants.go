package main

import "fmt"

// could be defined outside the function
const A1 int = 1

func main(){
	fmt.Println(A1)

	// could within the functions
	const A2 = 2
	fmt.Println(A2)

	// could be defined inside a block
	const (
		A3 int = 5
		A4 float32 = 0.9
		A5 = "builder"
		A6 = true
	)
	fmt.Println(A3, A4, A5, A6)
}