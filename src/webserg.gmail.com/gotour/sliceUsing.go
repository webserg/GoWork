package main

import "fmt"

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func main() {
	var buffer [256]byte
	for i := 0; i < len(buffer); i++ {
		fmt.Println(buffer[i])
	}
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println("after", buffer)
}
