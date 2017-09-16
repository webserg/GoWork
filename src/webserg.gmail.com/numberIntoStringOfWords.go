package webserg_gmail_com
// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println(os.Args[1])
	number := os.Args[1]
	fmt.Println(number)
	fmt.Println(len(number))
	sample := "12345"
	var sampleArray []string  = strings.Split(sample,"")
	fmt.Println(sampleArray)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
		fmt.Printf("%s ", sample[i])
	}

	singleDigits := [11]string{
		"zero",
		"one",
		"two",
		"thee",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	fmt.Println(singleDigits)


}
