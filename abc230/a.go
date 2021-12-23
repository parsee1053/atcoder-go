package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n >= 42 {
		n++
	}
	fmt.Printf("AGC%03d\n", n)
}
