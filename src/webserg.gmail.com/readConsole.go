package webserg_gmail_com
/*
8 9
1 4
1 2
4 2
4 3
3 2
5 2
3 5
8 2
8 6
*/
import (
	"fmt"
)

func main() {
	var n, m, x, y int
	fmt.Scan(&n)
	fmt.Scan(&m)
	edgesX := make([]int, m)
	edgesY := make([]int, m)
	for j := 0; j < m; j++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		edgesX[j] = x
		edgesY[j] = y
	}
	fmt.Print(n)
	fmt.Println(m)
	fmt.Println("edges:")
	for j := 0; j < m; j++ {
		fmt.Print(edgesX[j])
		fmt.Print(" ")
		fmt.Println(edgesY[j])
	}
	fmt.Println(edgesX)
}
