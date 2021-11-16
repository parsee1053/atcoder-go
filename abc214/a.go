package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := 8
	if n <= 125 {
		ans = 4
	} else if n <= 211 {
		ans = 6
	}
	fmt.Println(ans)
}
