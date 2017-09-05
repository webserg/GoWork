package main
import (
	"fmt"
	"time"
)
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
		fmt.Print(";r=", num)
	}
}

func ping(pings chan<- string, msg string) {
    pings <- msg
}
//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg + " pong"
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)

	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	go printCount(c)
	for _, v := range a {
		c <- v
		fmt.Print(";v = ", v)
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")

	pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)

}
