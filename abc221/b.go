package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if s == t {
		fmt.Println("Yes")
		return
	}
	sr := []rune(s)
	n := len(sr)
	ans := "No"
	for i := 0; i+1 < n; i++ {
		sr[i], sr[i+1] = sr[i+1], sr[i]
		if string(sr) == t {
			ans = "Yes"
		}
		sr[i], sr[i+1] = sr[i+1], sr[i]
	}
	fmt.Println(ans)
}
