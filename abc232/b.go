package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	sr := []rune(s)
	tr := []rune(t)
	n := len(sr)
	for k := 0; k <= 25; k++ {
		ok := true
		for i := 0; i < n; i++ {
			sc := (int(sr[i]-'a') + k) % 26
			tc := int(tr[i] - 'a')
			if sc != tc {
				ok = false
			}
		}
		if ok {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
