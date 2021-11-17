package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)
	two := int64(1)
	k := 0
	for two*2 <= n {
		two *= 2
		k++
	}
	fmt.Println(k)
}
