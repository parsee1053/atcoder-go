package main

import (
	"fmt"
)

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)
	at := int64(0)
	bt := int64(0)
	base := 1
	for a > 0 {
		at += int64(base * (a % 10))
		a /= 10
		base *= k
	}
	base = 1
	for b > 0 {
		bt += int64(base * (b % 10))
		b /= 10
		base *= k
	}
	ans := at * bt
	fmt.Println(ans)
}
