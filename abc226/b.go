package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	mp := map[string]bool{}
	for i := 0; i < n; i++ {
		l := nextInt()
		a := make([]string, l)
		for j := 0; j < l; j++ {
			a[j] = strconv.Itoa(nextInt())
		}
		mp[strings.Join(a, ",")] = true
	}
	fmt.Println(len(mp))
}
