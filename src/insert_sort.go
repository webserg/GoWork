package main

import (
	"fmt"
)

func main() {
	primes := [6]int{3, 2, 5, 7, 11, 13}
	fmt.Println(primes)
	for j := 1; j < len(primes); j++ {
		var key = primes[j]
		var i = j - 1
		for {
			if i < 0 || primes[i] < key {
				break
			}
			primes[i+1] = primes[i]
			i = i - 1
		}
		primes[i+1] = key
	}
	fmt.Println(primes)

}
