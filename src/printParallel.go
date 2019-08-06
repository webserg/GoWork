package main

import (
	"fmt"
	"time"
)

func getName() string {
	return "World!"
}

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}
func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)

	/*use channels*/
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	go printCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
