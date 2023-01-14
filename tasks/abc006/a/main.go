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

	ans := "NO"
	if n%3 == 0 || n == 3 {
		ans = "YES"
	}
	fmt.Fprintln(writer, ans)
}
