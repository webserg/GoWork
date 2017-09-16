package webserg_gmail_com

import (
	"fmt"
	"math"
)

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	return a + b
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() (ret int) {
		ret, a, b = a, b, a + b
		return
	}
}

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	p := func(a, b int) int {
		return a + b
	}

	fmt.Println(p(1, 2))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2 * i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

