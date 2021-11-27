package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	ans := "Yes"
	for i1 := 0; i1 < h; i1++ {
		for i2 := i1 + 1; i2 < h; i2++ {
			for j1 := 0; j1 < w; j1++ {
				for j2 := j1 + 1; j2 < w; j2++ {
					if a[i1][j1]+a[i2][j2] > a[i2][j1]+a[i1][j2] {
						ans = "No"
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
