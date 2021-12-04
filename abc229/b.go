package main

import (
	"fmt"
)

func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	ans := "Easy"
	for a > 0 && b > 0 {
		if (a%10)+(b%10) >= 10 {
			ans = "Hard"
		}
		a /= 10
		b /= 10
	}
	fmt.Println(ans)
}
