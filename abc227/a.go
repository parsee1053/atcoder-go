package main

import (
	"fmt"
)

func main() {
	var n, k, a int
	fmt.Scan(&n, &k, &a)
	ans := (a + k - 1) % n
	if ans == 0 {
		ans = n
	}
	fmt.Println(ans)
}
