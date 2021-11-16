package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)
	ans := "Strong"
	if x[0] == x[1] && x[1] == x[2] && x[2] == x[3] {
		ans = "Weak"
	}
	x0 := x[0] - '0'
	x1 := x[1] - '0'
	x2 := x[2] - '0'
	x3 := x[3] - '0'
	if (x0+1)%10 == x1 && (x1+1)%10 == x2 && (x2+1)%10 == x3 {
		ans = "Weak"
	}
	fmt.Println(ans)
}
