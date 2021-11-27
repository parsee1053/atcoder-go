package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	if s[n-2:] == "er" {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
