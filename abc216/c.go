package main

import (
	"fmt"
)

var ans = ""

func solve(n int64) {
	if n == 0 {
		return
	}
	if n%2 == 1 {
		ans = "A" + ans
		solve(n - 1)
	} else {
		ans = "B" + ans
		solve(n / 2)
	}
}

func main() {
	var n int64
	fmt.Scan(&n)
	solve(n)
	fmt.Println(ans)
}
