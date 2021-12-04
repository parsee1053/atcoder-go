package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	ans := "Yes"
	if s1 == "#." && s2 == ".#" {
		ans = "No"
	}
	if s1 == ".#" && s2 == "#." {
		ans = "No"
	}
	fmt.Println(ans)
}
