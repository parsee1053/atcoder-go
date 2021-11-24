package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	ans1 := s
	ans2 := s
	for i := 0; i < n; i++ {
		var t string
		t = s[1:n] + s[0:1]
		if t < ans1 {
			ans1 = t
		}
		if t > ans2 {
			ans2 = t
		}
		s = t
	}
	fmt.Println(ans1)
	fmt.Println(ans2)
}
