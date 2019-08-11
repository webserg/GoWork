package main

import (
	"fmt"
)

func main() {
	primes := []int{2, 3, 0, 1}
	fmt.Println(primes)
	merge(primes, 0, 1, 3)
	fmt.Println(primes)
}

func merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	//n2 := r - q
	L := make([]int, len(A[p:n1])+1)
	R := make([]int, len(A[n1:r+1])+1)
	copy(L, A[p:n1])
	copy(R, A[n1:r+1])
	i, j := 0, 0
	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}
