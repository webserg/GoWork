package main

import (
	"fmt"
	"math"
)

func main() {
	primes := []int{5, 2, 4, 7, 1, 3, 2}
	//primes := []int{5, 2, 4}
	fmt.Println(primes)
	//merge(primes, 0, 1, 2)
	mergeSort(primes, 0, len(primes)-1)
	fmt.Println(primes)
}

func mergeSort(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(A, p, q)
		mergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}

func merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}
	i, j := 0, 0
	L[n1] = math.MaxInt32
	R[n2] = math.MaxInt32
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}
