package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	ans := float64(d) / 100
	fmt.Printf("%.5f\n", ans)
}
