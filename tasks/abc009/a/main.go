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

	var N int
	fmt.Fscan(reader, &N)

	ans := N / 2
	if N%2 == 1 {
		ans++
	}
	fmt.Fprintln(writer, ans)
}
