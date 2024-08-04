package main

import "fmt"

func main(){
	var i,j = "hello", "world"
	fmt.Print(i)
  	fmt.Print(j, "\n")
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	fmt.Print(i, "\n",j)
	fmt.Print(i, " ", j)
	fmt.Println(i,j)

	var m,k = 10,20
	fmt.Print(m,k, "\n")

	// string formatting
	var o string = "Hello"
  	var p int = 15

	fmt.Printf("i has value: %v and type: %T\n", o, o)
	fmt.Printf("j has value: %v and type: %T", p, p)
	
	var n = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", n)
	fmt.Printf("%#v\n", n)
	fmt.Printf("%v%%\n", n)
	fmt.Printf("%T\n", n)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt) 

	var x = 15
	fmt.Printf("%b\n", x)
	fmt.Printf("%d\n", x)
	fmt.Printf("%+d\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%O\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%X\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%4d\n", x)
	fmt.Printf("%-4d\n", x)
	fmt.Printf("%04d\n", x)


	var txt2 = "Hello"
	fmt.Printf("%s\n", txt2)
	fmt.Printf("%q\n", txt2)
	fmt.Printf("%8s\n", txt2)
	fmt.Printf("%-8s\n", txt2)
	fmt.Printf("%x\n", txt2)
	fmt.Printf("% x\n", txt2)


	var t = true
	var f = false
	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)

	var pi = 3.141
	fmt.Printf("%e\n", pi)
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)
	fmt.Printf("%6.2f\n", pi)
	fmt.Printf("%g\n", pi)
}