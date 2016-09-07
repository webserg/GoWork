package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 10
	step := func() float64 {
		return z - (z * z - x) / (2 * z)
	}
	for zz := step(); math.Abs(zz - z) > 0.0001
	{
		z = zz
		zz = step()
	}
	return

}

func main() {
	fmt.Println(Sqrt(500))
	fmt.Println(math.Sqrt(500))
}
