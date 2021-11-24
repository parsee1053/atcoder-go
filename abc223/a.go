package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	ans := "No"
	if x != 0 && x%100 == 0 {
		ans = "Yes"
	}
	fmt.Println(ans)
}
