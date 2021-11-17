package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	t := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &t[i])
	}
	ok := false
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && t[i] == t[j] {
				ok = true
			}
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
