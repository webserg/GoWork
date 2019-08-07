package main

import (
	"container/list"
	"fmt"

	"unicode"
	"strings"
	"strconv"
	"math/rand" 
)

func main() {



	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	fmt.Println()
	ss := strings.FieldsFunc("  foo1;bar2,baz3...", f)
	counter:=0
	var l list.List
	l.PushBack(ss)
	for _,v := range ss {
		l.PushBack(v)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		counter++
	}
	fmt.Printf(strconv.Itoa(counter))
	fmt.Println("hfghf ")
	array := make([]byte, 10)
    rand.Read(array)
    fmt.Println(array)

	for i := 1; i < len(array); i += 4 {
    	v := array[i]	
    	fmt.Println(v)
	}
}
