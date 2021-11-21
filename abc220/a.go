package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	ans := -1
	for i := a; i <= b; i++ {
		if i%c == 0 {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
