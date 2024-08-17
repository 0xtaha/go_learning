package main

import "fmt"

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func main() {

	sum2 := 0
	for i := 0; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum2 += i
	}
	fmt.Println(sum2) // 6 (2+4)

	fmt.Println(unhex(2))
}
