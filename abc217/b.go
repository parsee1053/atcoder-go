package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]int)
	for i := 0; i < 3; i++ {
		var s string
		fmt.Scan(&s)
		mp[s]++
	}
	if mp["ABC"] == 0 {
		fmt.Println("ABC")
	}
	if mp["ARC"] == 0 {
		fmt.Println("ARC")
	}
	if mp["AGC"] == 0 {
		fmt.Println("AGC")
	}
	if mp["AHC"] == 0 {
		fmt.Println("AHC")
	}
}
