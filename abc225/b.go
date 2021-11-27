package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := make([]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		cnt[a-1]++
		cnt[b-1]++
	}
	ans := "No"
	for i := 0; i < n; i++ {
		if cnt[i] == n-1 {
			ans = "Yes"
		}
	}
	fmt.Println(ans)
}
