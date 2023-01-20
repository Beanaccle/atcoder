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

	var a string
	fmt.Fscan(reader, &a)

	ans := "a"
	if a == "a" {
		ans = "-1"
	}
	fmt.Fprintln(writer, ans)
}
