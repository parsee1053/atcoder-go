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
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		p := nextInt()
		ans[p-1] = strconv.Itoa(i + 1)
	}
	fmt.Println(strings.Join(ans, " "))
}
