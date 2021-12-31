package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	mp := make(map[string]int)
	ans := ""
	cnt := 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		mp[s]++
		if mp[s] > cnt {
			cnt = mp[s]
			ans = s
		}
	}
	fmt.Println(ans)
}
