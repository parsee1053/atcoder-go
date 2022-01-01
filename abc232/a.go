package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	a := int(rune(s[0]) - '0')
	b := int(rune(s[2]) - '0')
	fmt.Println(a * b)
}
