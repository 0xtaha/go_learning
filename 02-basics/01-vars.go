package main

import "fmt"

func main(){
	// variabels creation
	var a int = 1
	var b float32 = 0.5
	var c string = "fight"
	var d bool = true
	fmt.Println(a,b,c,d)

	// variabels creation with assigning directly
	// Can only be used inside functions
	a2:= 1
	b2:= 0.3
	c2:= "back"
	d2:= false
	fmt.Println(a2,b2,c2,d2)


	// create without assigning
	var a3 int //0
	var b3 float32 //0.0
	var c3 string //""
	var d3 bool //False
	fmt.Println(a3,b3,c3,d3)


	// assigining in a bulk
	var a4, b4 = 6, "Hello"
  	c4, d4 := 7, "World!"
	fmt.Println(a4,b4,c4,d4)

	
	var (
		a5 int
		b5 float32 = 1.3
		c5 string = "hello"
		d5 bool = false
	  )
	  fmt.Println(a5,b5,c5,d5)

	// variables are case senstive
	var age = 12
	var Age = 15
	fmt.Println(age)
	fmt.Println(Age)
}