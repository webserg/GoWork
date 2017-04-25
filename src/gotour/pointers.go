package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
}

var (
	v1 = Vertex{1, 2, "some"}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2,"some"} // has type *Vertex
)

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2, "hello"}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)

	pp := &v//To access the field X of a struct when we have the struct pointer p 
	pp.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
