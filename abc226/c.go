package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	t := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = nextInt()
		k := nextInt()
		a[i] = make([]int, k)
		for j := 0; j < k; j++ {
			a[i][j] = nextInt()
			a[i][j]--
		}
	}
	ans := int64(0)
	use := make([]bool, n)
	use[n-1] = true
	for i := n - 1; i >= 0; i-- {
		if use[i] {
			ans += int64(t[i])
			len := len(a[i])
			for j := 0; j < len; j++ {
				if !use[a[i][j]] {
					use[a[i][j]] = true
				}
			}
		}
	}
	fmt.Println(ans)
}
