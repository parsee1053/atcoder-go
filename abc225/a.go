package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if s[0] == s[1] && s[1] == s[2] {
		fmt.Println(1)
	} else if s[0] != s[1] && s[1] != s[2] && s[2] != s[0] {
		fmt.Println(6)
	} else {
		fmt.Println(3)
	}
}
