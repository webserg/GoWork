package main

import (
	"fmt"
)

func main() {
	var n, m, x, y int
	fmt.Scan(&n)
	fmt.Scan(&m)
	edgesX := make([]int, m)
	edgesY := make([]int, m)
	for j := 0; j < m; j++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		edgesX[j] = x
		edgesY[j] = y
	}
	fmt.Print(n)
	fmt.Println(m)
	fmt.Println("edges:")
	for j := 0; j < m; j++ {
		fmt.Print(edgesX[j])
		fmt.Print(" ")
		fmt.Println(edgesY[j])
	}
}
