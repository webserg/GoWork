package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)
//Sqrt make sqrt
func Sqrt(x float64) (z float64) {
	z = 10
	step := func() float64 {
		return z - (z * z - x) / (2 * z)
	}
	for zz := step(); math.Abs(zz - z) > 0.0001
	{
		z = zz
		zz = step()
		fmt.Println(zz)
	}
	return

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%g\n",v)
	}
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%g\n",v)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func deferUsing(){
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	fmt.Println(Sqrt(500))
	fmt.Println(math.Sqrt(500))
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	deferUsing()
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
