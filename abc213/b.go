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

type data struct {
	score int
	index int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]data, n)
	for i := 0; i < n; i++ {
		a[i].score = nextInt()
		a[i].index = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].score < a[j].score })
	fmt.Println(a[n-2].index)
}
