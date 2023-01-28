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

	cnt := 0
	for i := 0; i < N; i++ {
		var S string
		fmt.Fscan(reader, &S)

		if S == "For" {
			cnt++
		}
	}

	ans := "Yes"
	if (N / 2) >= cnt {
		ans = "No"
	}
	fmt.Fprintln(writer, ans)
}
