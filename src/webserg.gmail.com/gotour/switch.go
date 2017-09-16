package main

import (
	"fmt"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}
/**
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	var i, j int = 1, 2
	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	k := 3
	i = 2
	fmt.Println(i, j, k, c, python, java)
}
