package main

import (
	"fmt"
	"strconv"
)

func main() {
	mp := make(map[string]string)
	for i := 1; i <= 3; i++ {
		var s string
		fmt.Scan(&s)
		mp[strconv.Itoa(i)] = s
	}
	var t string
	fmt.Scan(&t)
	for i := 0; i < len(t); i++ {
		fmt.Print(mp[string(t[i])])
	}
	fmt.Println()
}
