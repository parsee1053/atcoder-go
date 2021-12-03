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
	x := nextInt()
	x--
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		a[i]--
	}
	used := make([]bool, n)
	ans := 0
	for {
		if used[x] {
			break
		}
		ans++
		used[x] = true
		x = a[x]
	}
	fmt.Println(ans)
}
