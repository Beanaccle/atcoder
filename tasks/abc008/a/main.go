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

	var s, t int
	fmt.Fscan(reader, &s, &t)

	fmt.Fprintln(writer, t-s+1)
}
