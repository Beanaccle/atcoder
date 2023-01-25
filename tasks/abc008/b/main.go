package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	votes := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		votes[s]++
	}

	max := 0
	ans := ""
	for k, v := range votes {
		if v > max {
			max = v
			ans = k
		}
	}

	fmt.Fprintln(writer, ans)
}
