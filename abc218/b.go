package main

import (
	"fmt"
)

func main() {
	p := make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Scan(&p[i])
	}
	ans := ""
	for i := 0; i < 26; i++ {
		ans += string('a' - 1 + p[i])
	}
	fmt.Println(ans)
}
