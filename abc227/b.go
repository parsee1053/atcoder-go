package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	ans := n
	for i := 0; i < n; i++ {
		ok := false
		for a := 1; a <= s[i]; a++ {
			for b := 1; b <= s[i]; b++ {
				if 4*a*b+3*(a+b) == s[i] {
					ok = true
				}
			}
		}
		if ok {
			ans--
		}
	}
	fmt.Println(ans)
}
