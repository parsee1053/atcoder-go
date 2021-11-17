package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	if s[n-1] <= '2' {
		fmt.Println(s[:n-2] + "-")
	} else if s[n-1] >= '7' {
		fmt.Println(s[:n-2] + "+")
	} else {
		fmt.Println(s[:n-2])
	}
}
