package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := "Alloy"
	if a == 0 {
		ans = "Silver"
	} else if b == 0 {
		ans = "Gold"
	}
	fmt.Println(ans)
}
