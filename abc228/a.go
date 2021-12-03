package main

import (
	"fmt"
)

func main() {
	var s, t, x int
	fmt.Scan(&s, &t, &x)
	ans := "No"
	if s < t {
		if s <= x && x < t {
			ans = "Yes"
		}
	} else {
		if x < t || s <= x {
			ans = "Yes"
		}
	}
	fmt.Println(ans)
}
