package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var t string
	for i := 0; i < 100; i++ {
		t += "oxx"
	}
	n := len(s)
	ans := "No"
	for i := 0; i < 3; i++ {
		if s == t[i:i+n] {
			ans = "Yes"
		}
	}
	fmt.Println(ans)
}
