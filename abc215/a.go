package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := "WA"
	if s == "Hello,World!" {
		ans = "AC"
	}
	fmt.Println(ans)
}
