package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	d := a - b
	ans := 1
	for i := 0; i < d; i++ {
		ans *= 32
	}
	fmt.Println(ans)
}
