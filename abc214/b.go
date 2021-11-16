package main

import (
	"fmt"
)

func main() {
	var s, t int
	fmt.Scan(&s, &t)
	ans := 0
	for a := 0; a <= s; a++ {
		for b := 0; b <= s; b++ {
			for c := 0; c <= s; c++ {
				if a+b+c <= s && a*b*c <= t {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
