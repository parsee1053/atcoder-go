package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type pair struct {
	value int
	index int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < m; i++ {
		b[i] = nextInt()
	}
	var c []pair
	for i := 0; i < n; i++ {
		c = append(c, pair{value: a[i], index: 0})
	}
	for i := 0; i < m; i++ {
		c = append(c, pair{value: b[i], index: 1})
	}
	sort.Slice(c, func(i, j int) bool { return c[i].value < c[j].value })
	ln := len(c)
	ans := 1010101010
	for i := 0; i+1 < ln; i++ {
		if c[i].index != c[i+1].index {
			diff := c[i+1].value - c[i].value
			if diff < ans {
				ans = diff
			}
		}
	}
	fmt.Println(ans)
}
