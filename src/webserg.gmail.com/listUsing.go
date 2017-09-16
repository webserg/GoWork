package webserg_gmail_com

import (
	"container/list"
	"fmt"

	"unicode"
	"strings"
	"strconv"
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
}
