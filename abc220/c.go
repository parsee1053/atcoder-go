package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var x int64
	fmt.Scan(&x)
	sum := int64(0)
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	t := x / sum
	ans := t * int64(n)
	total := t * sum
	i := 0
	for total <= x {
		total += a[i]
		i++
		ans++
	}
	fmt.Println(ans)
}
