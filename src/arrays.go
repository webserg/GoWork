package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slices
	var s []int = primes[1:4]
	fmt.Println(s[0])
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb)

	bb[0] = "XXX"
	fmt.Println(aa, bb)
	fmt.Println(names)

	str := []string{"this", "is", "a", "joined", "string\n"};
	fmt.Printf(strings.Join(str, " "));
	//4.1.2 Multi-dimensional arrays
	var multiArray [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multiArray[i][j] = fmt.Sprintf("row %d - column %d", i + 1, j + 1)
		}
	}
	fmt.Printf("%q", multiArray)
}
